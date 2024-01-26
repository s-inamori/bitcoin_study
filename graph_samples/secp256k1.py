import numpy as np
import matplotlib.pyplot as plt

def main():
    # https://stackoverflow.com/a/19757132
    _, ax = plt.subplots()
    ax.axhline(y=0, color='k')
    ax.axvline(x=0, color='k')
    b = 7
    y, x = np.ogrid[-5:5:100j, -5:5:100j]
    plt.contour(x.ravel(), y.ravel(), pow(y, 2) - pow(x, 3) - b, [0])
    plt.show()

if __name__ == '__main__':
    main()