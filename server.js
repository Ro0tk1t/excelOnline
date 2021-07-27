const Koa = require('koa');
const app = new Koa();

import Spreadsheet from 'x-data-spreadsheet';
import zhCN from 'x-data-spreadsheet/dist/locale/zh-cn';

Spreadsheet.locale('zh-cn', zhCN);
new Spreadsheet(document.getElementById('xss-demo'))

app.use(async ctx => {
    ctx.body = 'test'
});

app.listen(3000);
