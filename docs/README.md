# 基于wails的bilibili直播弹幕机

## Plan
- [x] 观众信息获取逻辑优化
- [x] Super Chat功能实现(一些细节尚待完善)
- [x] 切换直播间
- [x] 实现开箱即用无需配置的流程
- [ ] 礼物及上舰提醒
- [ ] 代码重构
- [ ] 响应式界面

## Usage
开箱即用，默认使用sqlite3在本地持久化数据，主要用于加快弹幕头像的获取并避免触发反爬机制

### Optional
在`config.yaml`或`config.yml`中修改文件内容进行可选项配置
1. 改用`postgreSQL`数据库以提高后端性能（需要另外运行数据库服务）
```yaml
database:
    type: postgres
    name: bliveDB
    host: 127.0.0.1
    port: 5432
    user: postgres
    password: 123456
```
修改`database`字段中`type`为`postgres`，并填写好数据库连接的相关配置

2. 使用代理
```yaml
proxy: http://127.0.0.1:xxxx
```
添加一级字段`proxy`

