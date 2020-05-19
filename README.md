# objectStorage
分布式对象存储

## 传统的网络存储

### NAS

NAS(Network Attached Storage)，是提供了一个存储功能和文件系统的网络服务器。客户端可以访问NAS上的文件系统，还可以上传和下载文件。对于客户端NAS就是网络文件服务器。

### SAN

SAN(Storage Area Network)，相对于NAS它只提供了块存储，文件系统的抽象交给客户端来管理。对于客户端来说，SAN就是一块磁盘，可以对其格式化、创建文件系统并挂载。

## 对象存储

### 优势

- 对象存储提高了存储系统的扩展性
- 能够以更低廉的成本提供数据的冗余能力


### Version(tag)

- 单机版本
- 可扩展分布式版本
- 元数据服务
- 数据校验和去重
- 数据冗余和及时修复
- 断点续传
- 数据压缩
- 数据维护 
