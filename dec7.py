#!/usr/bin/env python
import re


class ValueNode(object):
    ALL_ONES = 65535

    def __init__(self, nodes, lhs1, rhs):
        self.lhs1 = lhs1
        self.rhs = rhs
        self.nodes = nodes

    def value_from(self, name_or_number):
        if name_or_number.isdigit():
            return (int(name_or_number) & self.ALL_ONES)
        else:
            return self.nodes.get_value(name_or_number)

    def calculate(self):
        return self.value_from(self.lhs1)


class NotNode(ValueNode):
    def calculate(self):
        value = super(NotNode, self).calculate()
        return ~value & self.ALL_ONES


class BinaryNode(ValueNode):
    def __init__(self, nodes, lhs1, lhs2, rhs):
        super(BinaryNode, self).__init__(nodes, lhs1, rhs)
        self.lhs2 = lhs2

    def operation(self, value1, value2):
        raise ValueError('{0}, {1}'.format(value1, value2))

    def calculate(self):
        return self.operation(
            self.value_from(self.lhs1),
            self.value_from(self.lhs2)
        ) & self.ALL_ONES


class AndNode(BinaryNode):
    def operation(self, value1, value2):
        return (value1 & value2)


class OrNode(BinaryNode):
    def operation(self, value1, value2):
        return (value1 | value2)


class RshiftNode(BinaryNode):
    def operation(self, value1, value2):
        return (value1 >> value2)


class LshiftNode(BinaryNode):
    def operation(self, value1, value2):
        return (value1 << value2)


class NodeList(object):

    EXPRESSION_RE = re.compile(
        r"(?:(?P<arg1>\w+)\s+)(?:(?P<arg2>\w+)\s+)?(?:(?P<arg3>\w+)\s+)?->\s+(?P<result>\w+)"
    )
    def __init__(self, lines):
        self.names = {}
        self.values = {}
        for line in lines:
            matches = self.EXPRESSION_RE.match(line.strip())
            if matches:
                result = matches.group('result')
                arg1 = matches.group('arg1')
                arg2 = matches.group('arg2')
                arg3 = matches.group('arg3')
                if arg1 == 'NOT':
                    self.names[result] = NotNode(self, arg2, result)
                elif arg2 == 'AND':
                    self.names[result] = AndNode(self, arg1, arg3, result)
                elif arg2 == 'OR':
                    self.names[result] = OrNode(self, arg1, arg3, result)
                elif arg2 == 'LSHIFT':
                    self.names[result] = LshiftNode(self, arg1, arg3, result)
                elif arg2 == 'RSHIFT':
                    self.names[result] = RshiftNode(self, arg1, arg3, result)
                else:
                    self.names[result] = ValueNode(self, arg1, result)

    def get_value(self, name):
        if name not in self.names:
            raise ValueError

        if name in self.values:
            return self.values[name]

        value = self.names[name].calculate()
        self.values[name] = value
        return value

    def all_names(self):
        return self.names.keys()


if __name__ == '__main__':
    with open('input/day_7') as lines:
        nodes = NodeList(lines)

    if 'a' in nodes.all_names():
        print("a: {value}".format(value=nodes.get_value('a')))
