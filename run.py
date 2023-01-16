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
    for e in envs:
        if e in x:
            one['env'] = envs.get(e)
            break
    with open(filename, 'r') as f:
        data = json.load(f)
        one['title'] = data.get('info').get('title')
        one['version'] = data.get('info').get('version')
        f.close()

    config.append(one)

with open('./index-template.html') as template:
    content = template.read().replace("['ConfigPlaceholder']", json.dumps(config, ensure_ascii=False))
    template.close()
    with open("./page/index.html", 'w+') as index:
        index.write(content)
        index.close()
