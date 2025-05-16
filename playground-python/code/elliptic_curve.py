import matplotlib.pyplot as plt
import numpy as np

# Finite field modulus
p = 10000

# Curve parameters
a = 1
b = -1

O = None  # Point at infinity

# Modular inverse using Fermat's little theorem
def modinv(x):
    return pow(x, p - 2, p)

# Point addition over F_p
def ec_add(P, Q):
    if P is O:
        return Q
    if Q is O:
        return P
    if P[0] == Q[0] and (P[1] + Q[1]) % p == 0:
        return O
    if P == Q:
        m = (3 * P[0] * P[0] + a) * modinv(2 * P[1]) % p
    else:
        m = (Q[1] - P[1]) * modinv(Q[0] - P[0]) % p

    x_r = (m * m - P[0] - Q[0]) % p
    y_r = (m * (P[0] - x_r) - P[1]) % p
    return (x_r, y_r)

# Scalar multiplication using double-and-add
def ec_mul(k, P):
    result = O
    addend = P
    while k > 0:
        if k & 1:
            result = ec_add(result, addend)
        addend = ec_add(addend, addend)
        k >>= 1
    return result

# Find all curve points over F_p
points = []
for x in range(p):
    rhs = (x**3 + a * x + b) % p
    for y in range(p):
        if (y * y) % p == rhs:
            points.append((x, y))

# Select a known base point
P = (5, 1)

# Compute scalar multiples
multiples = []
for k in range(1, 20):
    Q = ec_mul(k, P)
    if Q is O:
        break
    multiples.append((Q[0], Q[1], k))

# Plot both finite field and real curve together
plt.figure(figsize=(10, 8))

# 1. Plot the real elliptic curve (over R) just for visualization
x_real = np.linspace(-3, 3, 1000000)
rhs_real = x_real**3 + a * x_real + b
y_real = np.sqrt(rhs_real, where=(rhs_real >= 0), out=np.full_like(rhs_real, np.nan))
plt.plot(x_real, y_real, label=r'$y = +\sqrt{x^3 + x - 1}$', color='gray', linestyle='dashed')
plt.plot(x_real, -y_real, label=r'$y = -\sqrt{x^3 + x - 1}$', color='gray', linestyle='dashed')

# 2. Plot the ECC points over F_p
for x in range(p):
    for y in range(p):
        plt.scatter(x, y, color='lightgray', s=10)

# Plot the actual curve points
if points:
    x_vals, y_vals = zip(*points)
    plt.scatter(x_vals, y_vals, color='blue', s=60, label='Curve Points $\mathbb{F}_{17}$')

# 3. Highlight scalar walk
for i, (x, y, k) in enumerate(multiples):
    plt.scatter(x, y, color='red', s=120)
    plt.text(x + 0.3, y + 0.3, f'{k}P', fontsize=9, color='black')
    if i > 0:
        x_prev, y_prev, _ = multiples[i - 1]
        plt.arrow(x_prev, y_prev, (x - x_prev), (y - y_prev),
                  head_width=0.3, head_length=0.3, length_includes_head=True, color='green', alpha=0.7)

plt.title(f'Scalar Walk over $\\mathbb{{F}}_{{{p}}}$ with Real Curve Overlay')
plt.xlabel('x')
plt.ylabel('y')
plt.grid(True)
plt.legend()
plt.xticks(range(p))
plt.yticks(range(p))
plt.xlim(-1, p)
plt.ylim(-1, p)
plt.show()
