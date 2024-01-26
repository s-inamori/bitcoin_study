import numpy as np
import matplotlib.pyplot as plt

def main():
    _, ax = plt.subplots()
    ax.axhline(y=0, color='k')
    ax.axvline(x=0, color='k')
    x = np.linspace(-10, 10, 100)
    y = 2 * x + 2
    plt.plot(x, y)
    plt.show()

if __name__ == '__main__':
    main()