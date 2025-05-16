# ARCHITECTURE

This document outlines the idea behind the architecture and how the code will be structured.

## CHRONOS ENGINE

The chronos engine will be the engine of the system. It will handle the file io and the searching through the database based on a queryplan or something similar.

The chronos engine will write bytes to files, files will be split in pages based on a certain size. The index page in a database will be an outline of what is contained in the file.

