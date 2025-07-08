const express = require('express');
const path = require('path');
const fs = require('fs');
const argv = require('minimist')(process.argv.slice(2)); // 解析命令行参数
const winston = require('winston');
const { format, transports } = winston;
const { combine, printf } = format;

const app = express();
const PORT = process.env.PORT || argv.port || 3000;
const VALID_SECRET = process.env.SECRET || `${argv.secret}` || '123456';
const STATIC_DIR = path.join(__dirname, 'public');
const UPLOAD_DIR = path.join(STATIC_DIR, 'docs');
const logFormat = printf(({ level, message, timestamp }) => {
    return `${timestamp} [${level.toUpperCase().padStart(5)}]：${message}`;
});
// 创建 logger 实例
const logger = winston.createLogger({
    level: 'debug', // 输出级别
    format: combine(
        format.timestamp({
            format: () => {
                const date = new Date();
                const year = date.getFullYear();
                const month = String(date.getMonth() + 1).padStart(2, '0');
                const day = String(date.getDate()).padStart(2, '0');
                const hours = String(date.getHours()).padStart(2, '0');
                const minutes = String(date.getMinutes()).padStart(2, '0');
                const seconds = String(date.getSeconds()).padStart(2, '0');
                const milliseconds = String(date.getMilliseconds()).padStart(3, '0');

                return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}.${milliseconds}`;
            }
        }),
        logFormat
    ),
    transports: [
        new transports.Console() // 输出到控制台
    ]
});

// 确保 public/docs 存在
if (!fs.existsSync(UPLOAD_DIR)) {
    fs.mkdirSync(UPLOAD_DIR, { recursive: true });
}

// 配置 multer 存储
const multer = require('multer');
const storage = multer.diskStorage({
    destination: (req, file, cb) => {
        cb(null, UPLOAD_DIR); // 文件保存路径
    },
    filename: (req, file, cb) => {
        const uniqueSuffix = Date.now() + '-' + Math.round(Math.random() * 1E9);
        const ext = path.extname(file.originalname);
        cb(null, file.fieldname + '-' + uniqueSuffix + ext); // 重命名文件以避免重复
    }
});

const upload = multer({ storage });

// 提供静态文件服务
app.use(express.static(STATIC_DIR));

function updateUrlsFile() {
    const docsDir = path.join(__dirname, 'public', 'docs');
    const urlsFilePath = path.join(__dirname, 'public', 'urls.json');

    try {
        // 读取 docs 目录中的所有文件
        const files = fs.readdirSync(docsDir).filter(file => file.endsWith('.json'));

        // 构建新的 urls 数据
        const urls = [];

        for (const file of files) {
            const filePath = path.join(docsDir, file);
            let title = '';
            let version = '';

            try {
                const data = JSON.parse(fs.readFileSync(filePath, 'utf-8'));

                // Swagger/OpenAPI 文档格式判断
                if (data.swagger) { // Swagger 2.0
                    title = data.info?.title || file;
                    version = data.info?.version || '';
                } else if (data.openapi) { // OpenAPI 3.x
                    title = data.info?.title || file;
                    version = data.info?.version || '';
                } else {
                    title = file;
                    version = '';
                }

            } catch (err) {
                logger.warn(`无法解析 ${file} 的 JSON 内容`);
                title = file;
                version = '';
            }

            urls.push({
                url: `./docs/${file}`,
                name: `${title}-${version}`.trim()
            });
        }

        // 写入 urls.json 文件
        fs.writeFileSync(urlsFilePath, JSON.stringify(urls, null, 2), 'utf-8');
        logger.info('urls.json 更新成功');
    } catch (err) {
        logger.error('更新 urls.json 失败:', err);
    }
}

// 处理上传请求
app.post('/upload', upload.single('file'), (req, res) => {
    const { secret, filename } = req.body;

    if (!filename) {
        return res.status(400).send('缺少 filename 参数');
    }

    if (secret !== VALID_SECRET) {
        return res.status(403).send('无效的 secret');
    }

    if (!req.file) {
        return res.status(400).send('未上传文件');
    }

    fs.rename(req.file.path, path.join(UPLOAD_DIR, filename), (err) => {
        if (err) {
            return res.status(500).send('文件重命名失败');
        }
        res.send(`文件 ${filename} 上传成功`);
    });

    updateUrlsFile()
});

app.listen(PORT, () => {
    logger.info(`服务已启动，访问 http://localhost:${PORT}`);
});
