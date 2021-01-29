# up-rancher
# 1、配合acme自动续签https证书。
## 1、1 acme提供
    1、免费的https证书，三个月的有效期
    2、支持泛域名https证书的申请
    3、创建定时任务自动续签申请证书
https://github.com/acmesh-official/acme.sh
## 1.2、 up-rancher
    1、将acme续签下来的https证书更新到rancher上

# 配合定时任务
可以完美实现 https泛域名的自动续签更新到rancher上

