# encoding: utf-8
# @author: pfwu
# @Date: 2018-03-24 23:34


def f(x):
    def g(y):
        return x + y
    return g


if __name__ == '__main__':
    a = f(1)
    b = f(1)
    print(a(1))
    print(b(2))
    print(a(1) == b(1))
