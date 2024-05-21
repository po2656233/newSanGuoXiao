# coding=utf-8
import os
import re

number = 0
# 基础类型列表
basic_types = ['double', 'float', 'int32', 'int64', 'uint32', 'uint64', 'sint32', 'sint64',
               'fixed32', 'fixed64', 'sfixed32', 'sfixed64', 'bool', 'string', 'bytes', 'repeated']


def listFiles(dirPath):
    fileList = []
    for root, dirs, files in os.walk(dirPath):
        for fileObj in files:
            if fileObj.endswith(".proto"):
                fileList.append(os.path.join(root, fileObj))
    return fileList


def dealSub(s1):
    _ = s1
    global number
    number = number + 1
    return '= ' + str(number) + ';'


def deal(s):
    m1 = '='
    m2 = ';'
    patN = re.compile(m1 + '(.*?)' + m2)
    data = s.group()
    resultN = patN.findall(str(data))
    if 0 == len(resultN):
        return data
    str1 = resultN[0].replace(' ', resultN[0])
    if int(str1) == 0:
        return data
    global number
    number = 0
    word = re.sub(patN, dealSub, str(data))
    return word


def camel_case(name, lower='true'):
    parts = name.split('_')
    if 1 < len(parts):
        first = parts[0].upper()
        if lower == 'true':
            first = parts[0].lower()
        name = first + ''.join(x.title() for x in parts[1:])
    elif 1 < len(name):
        first = name[0].upper()
        if lower == 'true':
            first = name[0].lower()
        name = first + name[1:]
    return name


def capitalize_custom_types(proto_content):
    # 使用正则表达式匹配字段
    field_pattern = re.compile(r'\s*(\w+)\s+(\w+)\s*=\s*\d+;')
    # 使用正则表达式查找所有字段
    fields = field_pattern.findall(proto_content)
    # 遍历所有字段
    for field in fields:
        # 获取字段的类型
        field_type = field[0]
        # 如果字段类型不是基础类型，则将其首字母改为大写
        if field_type not in basic_types:
            capitalized_type = field_type[0].upper() + field_type[1:]
            proto_content = proto_content.replace(f"{field_type} {field[1]}", f"{capitalized_type} {field[1]}")
    return proto_content


def modify_fields(proto_file_content):
    # 使用正则表达式匹配proto文件中的字段名，排除enum中的字段
    field_pattern = re.compile(r'(\s+)([a-zA-Z0-9_]+)(\s*=\s*\d+;)')

    # 找到所有字段，并转换为驼峰命名
    def camel_case_field(match):
        space_before = match.group(1)
        field_name = match.group(2)
        rest = match.group(3)
        return space_before + camel_case(field_name) + rest

    # 替换所有匹配的字段名
    proto_file_content = field_pattern.sub(camel_case_field, proto_file_content)
    proto_file_content = capitalize_custom_types(proto_file_content)

    # 处理枚举部分
    # 正则表达式匹配所有枚举定义
    enum_definitions = re.findall(r'enum\s+\w+\s*{[^}]*}', proto_file_content)
    # 创建字典来存储原始字段名和修改后的字段名
    field_replacements = {}
    for enum_def in enum_definitions:
        fields = re.findall(r'\b\w+\b\s*=', enum_def)
        for field in fields:
            original = field.split('=')[0]
            capitalized = camel_case(original, 'false')
            field_replacements[original] = capitalized

    # 替换原始protobuf定义中的字段名
    modified_protobuf_definition = proto_file_content
    for original, capitalized in field_replacements.items():
        modified_protobuf_definition = modified_protobuf_definition.replace(original, capitalized)
    return modified_protobuf_definition


def transform_struct(content, struct='message'):
    # 使用正则表达式找到所有的message定义
    # 这个正则表达式会匹配message后面跟着的任意单词字符（包括下划线），直到遇到{
    pattern = fr'{struct} (\w+)(\s*{{)'
    matches = re.findall(pattern, content)

    # 遍历所有匹配项，并将首字母大写
    for match in matches:
        # match是一个元组，第一个元素是message名称，第二个元素是大括号前的空白字符（包括可能存在的空格）
        message_name = match[0]
        whitespace = match[1]
        capitalized = message_name[0].upper() + message_name[1:]
        old = f'{struct} {message_name}{whitespace}'
        new = f'{struct} {capitalized}{whitespace}'
        content = content.replace(old, new)
    return content


def transform_proto(proto_content):
    content = transform_struct(proto_content, 'message')
    content = transform_struct(content, 'enum')
    return modify_fields(content)


def main():
    fileDir = "../internal/protocol"
    fileList = listFiles(fileDir)
    # 关键字1,2(修改引号间的内容)
    w1 = '{'
    w2 = '}'

    for fileObj in fileList:
        # 调整协议编码
        f = open(fileObj, 'r+', encoding='utf-8')
        buff = f.read()
        buff = transform_proto(buff)
        pat = re.compile(w1 + '(.*?)' + w2, re.S)
        pat.findall(buff)
        newNUMS = re.sub(pat, deal, buff)

        # 包名本地化
        f.seek(0)
        f.truncate()
        all_the_lines = newNUMS.split('\n')
        for line in all_the_lines:
            if len(line) > 0:  # 当改行为空，表明已经读取到文件末尾，退出循环
                content = line.split(' ')  # 因为每行有三个TAB符号分开的数字，将它们分开
                if len(content) > 0 and content[0] == 'package':
                    f.write("package pb;\n")
                else:
                    f.write(line + '\n')
        f.close()


if __name__ == '__main__':
    main()
