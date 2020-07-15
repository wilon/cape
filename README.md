# cape

国内下载一些文件经常`6.4 KB/s - 208 KB，共 30.0 MB，还剩 80 分钟`，下载不下来令人抓狂，有些时候 VPN 也比较慢。

此项目通过 GitHub Action 下载外网文件到阿里云 OSS，再自行下载。

经测试 GitHub Action 带宽 `46.46 MiB p/s`，国内阿里云 OSS 访问速度也OK。

感谢微软爸爸。

### 使用

1. 开通阿里云 OSS，拿到accessKeyID、accessKeySecret、endpoint

2. `fork` 这个 repo

3. 这个 repo 下：Settings->Secrets，配置：ENDPOINT、ACCESS_KEY_ID、ACCESS_KEY_SECRET、OSSPATH、BUCKET_NAME

4. 修改`url`文件，放入想下载的链接

5. `push`，然后 Actions 看结果

### 阿里云 OSS 简易计费情况

- 存储 1G 不到 1毛钱
- 上传完全免费！多少流量都不怕
- 下载流量费：1G 5 毛钱
- 请求次数：1 万次 1 毛钱
- 图片处理：10T 以内免费

计费详情：https://www.aliyun.com/price/product?#/oss/detail