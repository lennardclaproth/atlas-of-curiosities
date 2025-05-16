# ChronosDB: Paging & PageWriter Architecture

This document outlines the internal paging mechanism of ChronosDB and how the `PageWriter` is responsible for efficiently writing events to disk using a paged file layout.

---

let's first outline what paging is in database programming.

Paging is a function of memory management where a computer will store and retrieve data from a device's secondary storage to the primary storage. Memory management is a crucial aspect of any computing device, paging specifically is important to the implementation of virtual memory.

Paging works by reading and writing data to a secondary storage system (this is the storage system that is a more long term storage but slower, HDD/SSD) and/or loading it into primary memory (short term storage/RAM).

In a memory management system that takes advantage of paging, the OS reads data from secondary storage in blocks called pages, most of the time in a multiple of 4 i.e. 4, 16, 32, which have identical sizes. The physical region of memory containing a page is called a frame. When paging is used, a frame does not have to comprise a single contiguous region in secondary storage.

## Paging System

ChronosDB organizes its data storage into fixed-size **pages** (typically 4–16 KB). Each page contains multiple variable-length records (events) and supports fast insertions and lookups using a **slotted page layout**.

### 🔹 Page Structure (Slotted Layout)

Each page is composed of:

- A **Page Header**: metadata about the page (type, free space, etc.)
- A **Slot Directory**: a list of pointers to the actual records
- **Record**: records stored from the bottom up

Each **slot entry** includes:

- Offset to the record within the page
- Length of the record
- Optional overflow page reference (if the record is too large)

Each **Record** consists of the following:

---

## Overflow Pages

If a record is too large to fit into a single page (e.g., a bulk update event), ChronosDB uses **overflow pages** to store the data in chunks.

Each overflow page contains:

- A pointer to the **next overflow page** (or zero if it’s the last)
- A **slice of the original record**

The slot in the original page points to the first overflow page.

---

## PageWriter

The `PageWriter` component is responsible for writing events to disk, allocating new pages as needed, and managing overflow scenarios.

### Key Features

- Accepts batches of `ChronosEvent`
- Packs multiple events into a page until full
- Automatically writes to **new pages** when the current page overflows
- Detects **oversized records** and writes them using **overflow page chains**
- Updates the **slot directory** with record metadata

### Record insertion

1. Serialize event to binary
2. Check if it fits into the current page
   - if yes: insert into next available slot
   - if no:
     - if it is a small: start new page and try again
     - if it is large: write to overflow pages
3. Mark the page as dirty
4. When page is full or flushed: write to disk.

***\* records get added in slot from the back of the page***

## Frame layout

Each Frame file consists of:

- Data Pages: slotted pages containing events
- Overflow Pages: chained pages for large records

The frame of the data also have a header file and an index file:

- Header Page (Page 0): database version, configuration
- Index Page (Page 1): B-tree root or stream directory

## Resources

1. [What is paging](https://www.techtarget.com/whatis/definition/paging#:~:text=Paging%20is%20a%20function%20of,the%20implementation%20of%20virtual%20memory.)
