import os
import re

# 定义目录和文件路径
PROTO_DIR = r'..\internal\protocol'  # 修改为您的 proto 文件目录
GO_FILE = r'..\nodes\game\module\player\handle_msg.go'  # 请根据实际路径修改

def extract_req_messages(proto_dir):
    """
    提取所有未被注释且以 Req 结尾的 message 及其注释。
    返回一个列表，每个元素为元组 (message_name, comment)
    """
    req_messages = []
    message_pattern = re.compile(r'^\s*message\s+(\w+Req)\s*\{')

    for root, _, files in os.walk(proto_dir):
        for file in files:
            if file.endswith('.proto'):
                path = os.path.join(root, file)
                try:
                    with open(path, 'r', encoding='utf-8') as f:
                        lines = f.readlines()
                        for i, line in enumerate(lines):
                            msg_match = message_pattern.match(line)
                            if msg_match:
                                msg_name = msg_match.group(1)
                                
                                # 检查当前行是否被单行注释 '//' 注释掉
                                if re.match(r'^\s*//', line):
                                    continue  # 跳过被注释的 message

                                # 收集上方的连续单行注释
                                comments = []
                                j = i - 1
                                while j >= 0:
                                    comment_match = re.match(r'^\s*//\s?(.*)', lines[j])
                                    if comment_match:
                                        comments.insert(0, comment_match.group(1).strip())
                                        j -= 1
                                    else:
                                        break
                                full_comment = '\n'.join(comments) if comments else ''
                                req_messages.append((msg_name, full_comment))
                except Exception as e:
                    print(f'无法读取文件 {path}: {e}')
    return req_messages

def generate_registration_lines(messages):
    """
    生成注册代码行，不包含注释。
    """
    lines = []
    for msg, _ in messages:
        method = msg[:-3]  # 去除 'Req' 后缀
        lines.append(f'\tp.Local().Register(p.{method})')
    return lines

def generate_handler_functions(messages):
    """
    生成处理函数代码，包含注释。
    """
    functions = []
    for msg, comment in messages:
        method = msg[:-3]  # 去除 'Req' 后缀
        if comment:
            comment_lines = '// ' + comment.replace('\n', '\n// ')
            functions.append(f'''
{comment_lines}
func (p *ActorPlayer) {method}(session *cproto.Session, m *protoMsg.{msg}) {{
\t// TODO: 实现 {method} 处理逻辑
}}
''')
        else:
            functions.append(f'''
func (p *ActorPlayer) {method}(session *cproto.Session, m *protoMsg.{msg}) {{
\t// TODO: 实现 {method} 处理逻辑
}}
''')
    return functions

def update_go_file(go_file, registration_lines, handler_functions):
    try:
        if not os.path.isfile(go_file):
            print(f'文件不存在: {go_file}')
            return

        with open(go_file, 'r', encoding='utf-8') as f:
            content = f.read()

        updated_content = content  # 初始化

        # 匹配 ActorPlayer 的 registerLocalMsg 方法
        register_func_pattern = re.compile(
            r'(func\s+\(p\s+\*ActorPlayer\)\s+registerLocalMsg\s*\(\)\s*\{)([^}]*)\}',
            re.MULTILINE | re.DOTALL
        )
        match = register_func_pattern.search(content)
        if match:
            func_start = match.group(1)
            existing_body = match.group(2).strip()

            # 将现有注册行分割为集合，避免重复注册
            existing_lines = set()
            for line in existing_body.split('\n'):
                line = line.strip()
                if line.startswith('p.Local().Register'):
                    existing_lines.add(line)

            new_lines = []
            for reg_line in registration_lines:
                if reg_line.strip() not in existing_lines:
                    new_lines.append(reg_line)

            if new_lines:
                # 添加新的注册行到现有方法体
                updated_body = existing_body + '\n' + '\n'.join(new_lines)
                updated_func = f'{func_start}\n{updated_body}\n}}'
                updated_content = content.replace(match.group(0), updated_func)
                print('已在现有的 registerLocalMsg 方法中添加注册代码。')
            else:
                print('所有注册消息已存在，未执行更新。')
        else:
            # 如果没有找到 registerLocalMsg 方法，则添加该方法
            func_definition = '\nfunc (p *ActorPlayer) registerLocalMsg() {\n' + '\n'.join(registration_lines) + '\n}\n'
            updated_content += func_definition
            print('已添加 registerLocalMsg 方法并注册消息。')

        # 处理处理函数，避免重复
        handler_pattern = re.compile(r'func\s+\(p\s+\*ActorPlayer\)\s+(\w+)\s*\(', re.MULTILINE)
        existing_handlers = set(re.findall(handler_pattern, content))

        new_functions = []
        for func in handler_functions:
            func_name_match = re.search(r'func\s+\(p\s+\*ActorPlayer\)\s+(\w+)\s*\(', func)
            if func_name_match:
                func_name = func_name_match.group(1)
                if func_name not in existing_handlers:
                    new_functions.append(func)

        if new_functions:
            updated_content += '\n' + '\n'.join(new_functions)
            print('已添加新的处理函数。')
        else:
            print('所有处理函数已存在，未执行添加。')

        # 写回文件
        with open(go_file, 'w', encoding='utf-8') as f:
            f.write(updated_content)
        print(f'已更新 {go_file} 文件。')

    except OSError as e:
        print(f'文件操作错误: {e}')
    except Exception as e:
        print(f'发生错误: {e}')

def main():
    messages = extract_req_messages(PROTO_DIR)
    unique_messages = list(set(messages))  # 去重
    if not unique_messages:
        print('未找到以 Req 结尾的有效 message。')
        return

    # 生成注册代码
    registration_lines = generate_registration_lines(unique_messages)

    # 生成处理函数代码
    handler_functions = generate_handler_functions(unique_messages)

    # 更新 Go 文件
    update_go_file(GO_FILE, registration_lines, handler_functions)

if __name__ == "__main__":
    main()