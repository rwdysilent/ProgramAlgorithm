# encoding: utf-8
# @author: pfwu
# @Date: 2018-03-24 23:34


# def f():
#     x = 1
#
#     def g(y):
#         return x + y
#
#     return g
#
#
# if __name__ == '__main__':
#     a = f()
#     print(a(1))
#     print(a(2))


lists = []

for i in range(3):
    def func(x):
        return x*i
    lists.append(func)


print(lists)

for f in lists:
    print(f(2))
