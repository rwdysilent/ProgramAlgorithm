# encoding: utf-8
# @author: pfwu
# @Date: 2018-04-27 14:26

# N 个面包，分两种吃法，一次吃 1 个或者一次吃 2 个，求所有吃法的序列组合并打印


def eat(n):
    if n == 1:
        return ['1']
    if n == 2:
        return ['11', '2']
    return list(set(list(set(map(lambda x: '1' + x, eat(n - 1)))) +
                list(set(map(lambda x: '2' + x, eat(n - 2)))) +
                list(set(map(lambda x: '11' + x, eat(n - 2))))))


if __name__ == '__main__':
    print(eat(3))

    print(list(set(map(lambda x: '1' + x, ['1', '2', '1']))))
    print(list(set(map(lambda x: '2' + x, eat(3 - 2)))))
