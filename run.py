import os
import json

envs = {'test': '测试', 'prev': '预发', 'prod': '生产'}
config = []
files = os.listdir('./docs')

for x in files:
    filename = './docs/{name}'.format(name=x)
    if not os.path.isfile(filename) or not x.endswith('.json'):
        continue
    one = {}
    np = "未知"
    for e in envs:
        if e in x:
            np = envs.get(e)
            break
    with open(filename, 'r') as f:
        data = json.load(f)
        one['name'] = '[{env}]{title}-{version}'.format(env=np, title=data.get('info').get('version'), version=data.get('info').get('title'))
        one['url'] = filename

        # "name": fmt.Sprintf("[%s]%s-%s", np, sj.Info.Title, sj.Info.Version),
        # "url": filename,

        f.close()

    config.append(one)

with open('./index-template.html') as template:
    content = template.read().replace("['ConfigPlaceholder']", json.dumps(config, ensure_ascii=False))
    template.close()
    with open("./page/index.html", 'w+') as index:
        index.write(content)
        index.close()
