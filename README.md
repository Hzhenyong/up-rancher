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

环境变量
AK 、SK  是rancher的授权
RANCHER_URL 是要更新的证书的api地址 eg:https://118.178.142.114:8443/v3/project/c-cpqs7:p-cz7nd/certificates/p-cz7nd:rancher

WEB1  要更新证书的域名 hzy2013.cn



token-zl7p8
b5ttfb9pbdq98dz89p7k8thhvqdbmln4kl92pmrp59rlgxdkqkvfnz 


