## <font color="DeepSkyBlue">lua</font>
### <font color="DeepSkyBlue">aliyun</font>
#### <font color="DeepSkyBlue">dds</font>
##### <font color="DeepSkyBlue">get_db_perf</font>
###### Language: lua
###### Describe:
Query the performance of Alibaba Cloud MongoDB database.  
Reference document https://help.aliyun.com/document_detail/468425.html?spm=a2c4g.468308.0.0.41275d7afvzwFU .
###### Params:
dbInstanceId, string, Alibaba Cloud MongoDB instance ID.  
keys, table(1d), List of performance metric names.  
startTime, string, Query start time.  
endTime, string, Query end time.  
opts, table(k/v), Optional parameters, available optional parameters are as follows:<br>NodeId: Mongos node ID or Shard node ID in a sharded cluster instance, can be used to query the performance of a single node, only available when the DBInstanceId parameter is a sharded cluster instance ID.<br>RoleId: Node role ID of a single-node instance or replica set instance, can be queried by calling the DescribeReplicaSetRole interface, only available when the DBInstanceId parameter is a single-node instance ID or replica set instance ID.<br>ReplicaSetRole: Node role of a single-node instance or replica set instance, two values: Primary (primary node), Secondary (secondary node), only available when the DBInstanceId parameter is a single-node instance ID or replica set instance ID, when the DBInstanceId parameter is a single-node instance, this parameter only supports Primary.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get_db_perf_batch</font>
###### Language: lua
###### Describe:
Batch query the performance of Alibaba Cloud MongoDB databases.  
Reference document https://help.aliyun.com/document_detail/468425.html?spm=a2c4g.468308.0.0.41275d7afvzwFU .
###### Params:
reqs, table(1d), Each element represents a request, type is table(k/v), parameters are the same as get_db_perf.  
parallel, number(int), Degree of parallelism (1-16).
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">list_db_instances</font>
###### Language: lua
###### Describe:
List Alibaba Cloud MongoDB instances.  
Reference document https://help.aliyun.com/document_detail/468348.html?spm=a2c4g.468425.0.0.4a897e2aCfOHu9 .
###### Params:
opts, table(k/v), Optional parameters, available optional parameters refer to the reference document in describe.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
#### <font color="DeepSkyBlue">dns</font>
##### <font color="DeepSkyBlue">list_domain_records</font>
###### Language: lua
###### Describe:
List Alibaba Cloud DNS domain resolution records.  
Reference document https://help.aliyun.com/document_detail/29776.html?spm=a2c4g.29751.0.0.1c8870a7pJPU0c .
###### Params:
domainName, string, Domain name.  
opts, table(k/v), table(k/v), optional parameters, available optional parameters refer to the reference document in describe.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">list_domains</font>
###### Language: lua
###### Describe:
List Alibaba Cloud DNS domains.  
Reference document https://help.aliyun.com/document_detail/29751.html?spm=a2c4g.29751.0.0.2d9c3a40LzPUWA .
###### Params:
opts, table(k/v), Optional parameters, available optional parameters refer to the reference document in describe.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
#### <font color="DeepSkyBlue">mse</font>
##### <font color="DeepSkyBlue">get_nacos_config</font>
###### Language: lua
###### Describe:
Get configuration information under the specified namespace from Alibaba Cloud Microservice Engine MSE Nacos.
###### Params:
instanceId, string, Microservice Engine MSE instance ID.  
namespaceId, string, Microservice Engine MSE namespace ID.  
group, string, Microservice Engine MSE Nacos group.  
dataId, string, Microservice Engine MSE Nacos dataId.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">list_engine_namespaces</font>
###### Language: lua
###### Describe:
List namespace information in Alibaba Cloud Microservice Engine MSE.
###### Params:
instanceId, string, Microservice Engine MSE instance ID.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">list_nacos_configs</font>
###### Language: lua
###### Describe:
List configuration information under the specified namespace from Alibaba Cloud Microservice Engine MSE Nacos.
###### Params:
instanceId, string, Microservice Engine MSE instance ID.  
namespaceId, string, Microservice Engine MSE namespace ID.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
#### <font color="DeepSkyBlue">open_dds</font>
##### Language: lua
##### Describe:
Open an Alibaba Cloud DDS database connection.
##### Params:
endpoint, string, Alibaba Cloud endpoint.  
accessKeyId, string, Alibaba Cloud accessKeyId.  
accessKeySecret, string, Alibaba Cloud accessKeySecret.  
securityToken, string, Alibaba Cloud securityToken.
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">open_dns</font>
##### Language: lua
##### Describe:
Open an Alibaba Cloud DNS connection.
##### Params:
endpoint, string, Alibaba Cloud endpoint.  
accessKeyId, string, Alibaba Cloud accessKeyId.  
accessKeySecret, string, Alibaba Cloud accessKeySecret.  
securityToken, string, Alibaba Cloud securityToken.
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">open_mse</font>
##### Language: lua
##### Describe:
Open an Alibaba Cloud Microservice Engine MSE connection.
##### Params:
endpoint, string, Alibaba Cloud endpoint.  
accessKeyId, string, Alibaba Cloud accessKeyId.  
accessKeySecret, string, Alibaba Cloud accessKeySecret.  
securityToken, string, Alibaba Cloud securityToken.
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">open_oss</font>
##### Language: lua
##### Describe:
Open an Alibaba Cloud OSS (Object Storage Service) connection.
##### Params:
endpoint, string, Alibaba Cloud endpoint.  
accessKeyId, string, Alibaba Cloud accessKeyId.  
accessKeySecret, string, Alibaba Cloud accessKeySecret.  
securityToken, string, Alibaba Cloud securityToken.
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">open_sls</font>
##### Language: lua
##### Describe:
Open an Alibaba Cloud SLS (Log Service) connection.
##### Params:
endpoint, string, Alibaba Cloud endpoint.  
accessKeyId, string, Alibaba Cloud accessKeyId.  
accessKeySecret, string, Alibaba Cloud accessKeySecret.  
securityToken, string, Alibaba Cloud securityToken.
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">oss</font>
##### <font color="DeepSkyBlue">copy_object</font>
###### Language: lua
###### Describe:
Copy an object in Alibaba Cloud OSS.
###### Params:
srcBucketName, string, Source bucket name.  
srcObjectKey, string, Source object key.  
dstBucketName, string, Destination bucket name.  
dstObjectKey, string, Destination object key.  
opts, table(k/v), Optional parameters (ignored for now).
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">delete_object</font>
###### Language: lua
###### Describe:
Delete an object from the specified bucket in Alibaba Cloud OSS.
###### Params:
bucketName, string, Bucket name.  
objectKey, string, Object key.  
opts, table(k/v), Optional parameters (ignored for now).
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get_object_to_file</font>
###### Language: lua
###### Describe:
Download an object from the specified bucket in Alibaba Cloud OSS.
###### Params:
bucketName, string, Bucket name.  
objectKey, string, Object key.  
filaPath, string, Local file path.  
opts, table(k/v), Optional parameters (ignored for now).
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">is_object_exist</font>
###### Language: lua
###### Describe:
Check whether an object exists in the specified bucket in Alibaba Cloud OSS.
###### Params:
bucketName, string, Bucket name.  
objectKey, string, Object key.  
opts, table(k/v), Optional parameters (ignored for now).
###### Returns:
exist, bool or nil, Returns true or false on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">list_buckets</font>
###### Language: lua
###### Describe:
List buckets in Alibaba Cloud OSS.
###### Params:
opts, table(k/v), Optional parameters (ignored for now).
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">list_objects</font>
###### Language: lua
###### Describe:
List objects in the specified bucket in Alibaba Cloud OSS.
###### Params:
bucketName, string, Bucket name.  
opts, table(k/v), Optional parameters (currently only support marker, prefix, delimiter, max_keys).
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">put_object_from_file</font>
###### Language: lua
###### Describe:
Upload a file to the specified bucket in Alibaba Cloud OSS.
###### Params:
bucketName, string, Bucket name.  
objectKey, string, Object key.  
filePath, string, Local file path.  
opts, table(k/v), Optional parameters (ignored for now).
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">sign_url</font>
###### Language: lua
###### Describe:
Generate a temporary signed URL for an object in the specified bucket in Alibaba Cloud OSS.
###### Params:
bucketName, string, Bucket name.  
objectKey, string, Object key.  
method, string, HTTP method.  
expiredInSec, number(int), Expiration time (seconds).  
opts, table(k/v), Optional parameters (ignored for now).
###### Returns:
url, string or nil, Returns the signed URL on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
#### <font color="DeepSkyBlue">sls</font>
##### <font color="DeepSkyBlue">consume_store</font>
###### Language: lua
###### Describe:
Consume logs stored in the specified project and store in Alibaba Cloud SLS (Log Service).
###### Params:
project, string, Alibaba Cloud Log Service project.  
store, string, Alibaba Cloud Log Service logstore.  
from, string, Consumption start time (YYYY-MM-DD HH:MM:SS).  
query, string, Filter SQL.  
callback, func(table), Callback function.
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">list_project</font>
###### Language: lua
###### Describe:
List projects in Alibaba Cloud SLS (Log Service).
###### Params:
None.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">list_store</font>
###### Language: lua
###### Describe:
List stores under the specified project in Alibaba Cloud SLS (Log Service).
###### Params:
project, string, Alibaba Cloud Log Service project.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">put_log</font>
###### Language: lua
###### Describe:
Write logs to the specified project and store in Alibaba Cloud SLS (Log Service).
###### Params:
project, string, Alibaba Cloud Log Service project.  
store, string, Alibaba Cloud Log Service logstore.  
log, table(k/v), Detailed data.  
tag, table(k/v), Tag parameters, optional.  
opt, table(k/v), Optional parameters (category, topic, source, machineUUID), optional.
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
### <font color="DeepSkyBlue">archives</font>
#### <font color="DeepSkyBlue">zip_compress</font>
##### Language: lua
##### Describe:
Compress a file or folder to ZIP format.
##### Params:
src, string, Path to the file or folder to be compressed.  
dst, string, Path to the compressed file.  
level, number(int), Compression level (0-9, higher value means better compression but slower speed, default 6, optional).
##### Returns:
err, string, Compression error message.
#### <font color="DeepSkyBlue">zip_compress_subdir</font>
##### Language: lua
##### Describe:
Compress subdirectories to ZIP format, each subdirectory is compressed into a ZIP file with the same name.
##### Params:
src, string, Root directory to be compressed.  
dst, string, Root directory for storing compressed files.  
level, number(int), Compression level (0-9, higher value means better compression but slower speed).  
parallel, number(int), Degree of parallelism.  
ignoreExist, bool, Ignore subdirectories that have already been compressed (default true, optional).
##### Returns:
err, string, Compression error message.
#### <font color="DeepSkyBlue">zip_decompress</font>
##### Language: lua
##### Describe:
Decompress a ZIP file.
##### Params:
src, string, Path to the file to be decompressed.  
dst, string, Decompression destination path.
##### Returns:
err, string, Decompression error message.
#### <font color="DeepSkyBlue">zstd_compress</font>
##### Language: lua
##### Describe:
Compress a file or folder to Zstandard format.
##### Params:
src, string, Path to the file or folder to be compressed.  
dst, string, Path to the compressed file.  
level, number(int), Compression level (-5-22, higher value means better compression but slower speed, default 3, optional).
##### Returns:
err, string, Compression error message.
#### <font color="DeepSkyBlue">zstd_compress_subdir</font>
##### Language: lua
##### Describe:
Compress subdirectories to Zstandard format, each subdirectory is compressed into a file with the same name.
##### Params:
src, string, Root directory to be compressed.  
dst, string, Root directory for storing compressed files.  
level, number(int), Compression level (-5-22, higher value means better compression but slower speed).  
parallel, number(int), Degree of parallelism.  
ignoreExist, bool, Ignore subdirectories that have already been compressed (default true, optional).
##### Returns:
err, string, Compression error message.
#### <font color="DeepSkyBlue">zstd_decompress</font>
##### Language: lua
##### Describe:
Decompress a Zstandard file.
##### Params:
src, string, Path to the file to be decompressed.  
dst, string, Decompression destination path.
##### Returns:
err, string, Decompression error message.
### <font color="DeepSkyBlue">aws</font>
#### <font color="DeepSkyBlue">open_s3</font>
##### Language: lua
##### Describe:
Open an AWS S3 (Simple Storage Service) connection.
##### Params:
region, string, Region.  
endpoint, string, Endpoint.  
accessKeyId, string, AccessKeyId.  
accessKeySecret, string, AccessKeySecret.  
securityToken, string, SecurityToken.
##### Returns:
client, table, Client instance.
##### Since: 0.1
#### <font color="DeepSkyBlue">s3</font>
##### <font color="DeepSkyBlue">delete</font>
###### Language: lua
###### Describe:
Delete an object.
###### Params:
bucket, string, Bucket.  
key, string, Key.
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">download</font>
###### Language: lua
###### Describe:
Download an object.
###### Params:
bucket, string, Bucket.  
key, string, Key.  
fn, string, Local file path.  
opts, table(k/v), Optional parameters (ignored for now).
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">exist</font>
###### Language: lua
###### Describe:
Check if an object exists.
###### Params:
bucket, string, Bucket.  
key, string, Key.
###### Returns:
res, bool or nil, Returns true or false on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">upload</font>
###### Language: lua
###### Describe:
Upload an object.
###### Params:
bucket, string, Bucket.  
key, string, Key.  
fn, string, Local file path.  
opts, table(k/v), Optional parameters (ignored for now).
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
### <font color="DeepSkyBlue">buffers</font>
#### <font color="DeepSkyBlue">buffer</font>
##### <font color="DeepSkyBlue">reset</font>
###### Language: lua
###### Describe:
Reset the buffer.
###### Params:
fn, string, File path.
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">save</font>
###### Language: lua
###### Describe:
Save the buffer to a file.
###### Params:
fn, string, File path.  
compress, bool, Whether to compress (gzip).
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">write</font>
###### Language: lua
###### Describe:
Write data to the buffer.
###### Params:
elements, bool/number/string, Data elements to be appended (variable arguments).
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">new</font>
##### Language: lua
##### Describe:
Create a new buffer.
##### Params:
None.
##### Returns:
buf, table, The created buffer.
### <font color="DeepSkyBlue">cache</font>
#### <font color="DeepSkyBlue">del</font>
##### Language: lua
##### Describe:
Delete a cache entry.
##### Params:
k, string, Cache key.
##### Returns:
v, any, The value before deletion, returns nil if not exists.
#### <font color="DeepSkyBlue">get</font>
##### Language: lua
##### Describe:
Get a cache entry.
##### Params:
k, string, Cache key.
##### Returns:
v, any, The retrieved value, returns nil if not exists.
#### <font color="DeepSkyBlue">getf</font>
##### Language: lua
##### Describe:
Get a cache entry, set it if not exists.
##### Params:
k, string, Cache key.  
f, function() (any, error), function to get the cache value if not exists.  
ttl, number(int), Cache TTL (seconds), optional, default 0 means persistent storage.
##### Returns:
v, any or nil, The retrieved value, returns value on success, returns nil on failure.  
err, nil or string, Returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">reset</font>
##### Language: lua
##### Describe:
Reset the cache.
##### Params:
None.
##### Returns:
None.
#### <font color="DeepSkyBlue">set</font>
##### Language: lua
##### Describe:
Set a cache entry.
##### Params:
k, string, Cache key.  
v, any, Cache value.  
ttl, number(int), Cache TTL (seconds), optional, default 0 means persistent storage.
##### Returns:
p, any, The previous value, returns nil if not exists.
### <font color="DeepSkyBlue">charts</font>
#### <font color="DeepSkyBlue">bar</font>
##### Language: lua
##### Describe:
Bar chart.
##### Params:
params.title, string, Page title.  
params.fields, table(1d), Table headers (e.g. {{name="Date",key="date"},{name="Open",key="open"},{name="Close",key="close"}}).  
params.items, table(1d), Data list (e.g. {{date="2024-05-08",open=12.16,close=12.91},...}).  
params.xkey, string, X-axis data key (e.g. date).
##### Returns:
res, string or nil, Returns the HTML page source string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">kline</font>
##### Language: lua
##### Describe:
K-line chart.
##### Params:
params.title, string, Page title.  
params.items, table(1d), Data list (e.g. {{date="2024-05-08",open=12.16,close=12.91,high=13.43,low=11.95,percent=7.02,volume=2637823},...}).
##### Returns:
res, string or nil, Returns the HTML page source string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">line</font>
##### Language: lua
##### Describe:
Line chart.
##### Params:
params.title, string, Page title.  
params.fields, table(1d), Table headers (e.g. {{name="Date",key="date"},{name="Open",key="open"},{name="Close",key="close"}}).  
params.items, table(1d), Data list (e.g. {{date="2024-05-08",open=12.16,close=12.91},...}).  
params.xkey, string, X-axis data key (e.g. date).  
params.smooth, bool, Whether to smooth the curve.
##### Returns:
res, string or nil, Returns the HTML page source string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">pie</font>
##### Language: lua
##### Describe:
Pie chart.
##### Params:
params.title, string, Page title.  
params.items, table(1d), Data list (e.g. {{name="Search Engine",value=1048},{name="Direct",value=735},{name="Email",value=580},...}).
##### Returns:
res, string or nil, Returns the HTML page source string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">sankey</font>
##### Language: lua
##### Describe:
Sankey diagram.
##### Params:
params.title, string, Page title.  
params.links, table(1d), Data list (e.g. {{source="FRF",target="Colomiers",value=357.8399963378906},...}).
##### Returns:
res, string or nil, Returns the HTML page source string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">scatter</font>
##### Language: lua
##### Describe:
Scatter chart.
##### Params:
params.title, string, Page title.  
params.items, table(1d), Data list (e.g. {{sname="1990",dname="Australia",x=28604,y=77,z=17096869,size=8.269672061212585},...}).
##### Returns:
res, string or nil, Returns the HTML page source string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">table</font>
##### Language: lua
##### Describe:
Table.
##### Params:
params.title, string, Page title.  
params.fields, table(1d), Table headers (e.g. {{name="Date",key="date"},{name="Open",key="open"},{name="Close",key="close"},{name="High",key="high"},{name="Low",key="low"},{name="Change",key="percent"},{name="Volume",key="volume"}}).  
params.items, table(1d), Data list (e.g. {{date="2024-05-08",open=12.16,close=12.91,high=13.43,low=11.95,percent=7.02,volume=2637823},...}).
##### Returns:
res, string or nil, Returns the HTML page source string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">tree</font>
##### Language: lua
##### Describe:
Tree chart.
##### Params:
params.title, string, Page title.  
params.tree, table, Data (e.g. {name="flare",value=100,children={{name="analytics",value=20,children={...},collapsed=true},...},collapsed=true}).
##### Returns:
res, string or nil, Returns the HTML page source string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
### <font color="DeepSkyBlue">cmds</font>
#### <font color="DeepSkyBlue">exec</font>
##### Language: lua
##### Describe:
Execute a command (blocks until completion, returns command output).  
Optional parameters are as follows.  
opts.dir, string, working directory for the command.  
opts.env, table(1d), environment variables, list of strings.  
opts.timeout, number(int), timeout in seconds.
##### Params:
name, string, Command.  
args, table(1d), List of arguments, each argument is a string.  
opts, table(k/v), Optional parameters.
##### Returns:
res, string or nil, Command output (note: res may not be nil when err is non-nil).  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">run</font>
##### Language: lua
##### Describe:
Execute a command (blocks until completion, if scan is provided, command output is captured during execution, otherwise ignored).  
Optional parameters are as follows.  
opts.dir, string, working directory for the command.  
opts.env, table(1d), environment variables, list of strings.  
opts.timeout, number(int), timeout in seconds.
##### Params:
name, string, Command.  
args, table(1d), List of arguments, each argument is a string.  
scan, function(line) or nil, Function to process command output lines.  
opts, table(k/v), Optional parameters.
##### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">start</font>
##### Language: lua
##### Describe:
Execute a command (non-blocking, returns immediately, command output is ignored).  
Optional parameters are as follows.  
opts.dir, string, working directory for the command.  
opts.env, table(1d), environment variables, list of strings.
##### Params:
name, string, Command.  
args, table(1d), List of arguments, each argument is a string.  
opts, table(k/v), Optional parameters.
##### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.
### <font color="DeepSkyBlue">config</font>
#### <font color="DeepSkyBlue">cc_del</font>
##### Language: lua
##### Describe:
Delete a configuration from the configuration center.
##### Params:
namespace, string, Configuration namespace.  
key, string, Configuration key.
##### Returns:
err, nil or string, Returns nil on success, returns error message string on failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">cc_get</font>
##### Language: lua
##### Describe:
Get a configuration from the configuration center.
##### Params:
namespace, string, Configuration namespace.  
key, string, Configuration key.
##### Returns:
res, string or nil, Returns the configuration string value on success, returns nil on failure.  
err, nil or string, Returns nil on success, returns error message string on failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">cc_has</font>
##### Language: lua
##### Describe:
Check if a configuration exists in the configuration center.
##### Params:
namespace, string, Configuration namespace.  
key, string, Configuration key.
##### Returns:
res, bool or nil, Returns true if exists, false otherwise on success, returns nil on failure.  
err, nil or string, Returns nil on success, returns error message string on failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">cc_set</font>
##### Language: lua
##### Describe:
Set a configuration in the configuration center.
##### Params:
namespace, string, Configuration namespace.  
key, string, Configuration key.  
value, string, Configuration value.
##### Returns:
err, nil or string, Returns nil on success, returns error message string on failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">exist</font>
##### Language: lua
##### Describe:
Check if a configuration exists.
##### Params:
key, string, Configuration key.
##### Returns:
res, bool, Returns true if exists, false otherwise.
##### Since: 0.1
#### <font color="DeepSkyBlue">get_bool</font>
##### Language: lua
##### Describe:
Get a configuration value as boolean.
##### Params:
key, string, Configuration key.  
def, bool, Default value returned if the configuration does not exist.
##### Returns:
res, bool, The retrieved configuration value.
##### Since: 0.1
#### <font color="DeepSkyBlue">get_by_pattern</font>
##### Language: lua
##### Describe:
Get configurations by matching the key pattern using regex.
##### Params:
pattern, string, Regex pattern to match configuration keys.
##### Returns:
res, table(k/v), The retrieved configuration key-value pairs.
##### Since: 0.1
#### <font color="DeepSkyBlue">get_by_prefix</font>
##### Language: lua
##### Describe:
Get configurations by key prefix.
##### Params:
prefix, string, Configuration key prefix.
##### Returns:
res, table(k/v), The retrieved configuration key-value pairs.
##### Since: 0.1
#### <font color="DeepSkyBlue">get_duration</font>
##### Language: lua
##### Describe:
Get a configuration value as duration in seconds (int64).
##### Params:
key, string, Configuration key.  
def, number(int64), Default duration value in seconds returned if the configuration does not exist.
##### Returns:
res, string, The retrieved configuration value.
##### Since: 0.1
#### <font color="DeepSkyBlue">get_float</font>
##### Language: lua
##### Describe:
Get a configuration value as float64.
##### Params:
key, string, Configuration key.  
def, number(float64), Default value returned if the configuration does not exist.
##### Returns:
res, number(float64), The retrieved configuration value.
##### Since: 0.1
#### <font color="DeepSkyBlue">get_int</font>
##### Language: lua
##### Describe:
Get a configuration value as int.
##### Params:
key, string, Configuration key.  
def, number(int), Default value returned if the configuration does not exist.
##### Returns:
res, number(int), The retrieved configuration value.
##### Since: 0.1
#### <font color="DeepSkyBlue">get_slice</font>
##### Language: lua
##### Describe:
Get a configuration value as a table(1d) of strings.
##### Params:
key, string, Configuration key.  
def, table(1d), Default value returned if the configuration does not exist, elements are strings.
##### Returns:
res, table(1d), The retrieved configuration value, elements are strings.
##### Since: 0.1
#### <font color="DeepSkyBlue">get_string</font>
##### Language: lua
##### Describe:
Get a configuration value as string.
##### Params:
key, string, Configuration key.  
def, string, Default value returned if the configuration does not exist.
##### Returns:
res, string, The retrieved configuration value.
##### Since: 0.1
#### <font color="DeepSkyBlue">load</font>
##### Language: lua
##### Describe:
Load a custom .cfg configuration file.
##### Params:
fn, 配置文件路径, If the path is relative, it is relative to the working directory.
##### Returns:
res, table(k/v) or nil, Returns configuration data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">load_lst</font>
##### Language: lua
##### Describe:
Load a custom .lst configuration file.
##### Params:
fn, 配置文件路径, If the path is relative, it is relative to the working directory.
##### Returns:
res, table(1d) or nil, Returns configuration data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### Since: 0.1
### <font color="DeepSkyBlue">cryptos</font>
#### <font color="DeepSkyBlue">aes_decrypt</font>
##### Language: lua
##### Describe:
AES decryption.
##### Params:
content, string, String to be decrypted.  
key, string, Key.
##### Returns:
res, string or nil, Returns the decrypted string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">aes_encrypt</font>
##### Language: lua
##### Describe:
AES encryption.
##### Params:
content, string, String to be encrypted.  
key, string, Key.
##### Returns:
res, string or nil, Returns the encrypted string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">md5_sum</font>
##### Language: lua
##### Describe:
Calculate the MD5 checksum of a string.
##### Params:
content, string, String to be hashed.
##### Returns:
res, string, Hexadecimal checksum string.
#### <font color="DeepSkyBlue">md5_sum_file</font>
##### Language: lua
##### Describe:
Calculate the MD5 checksum of a file.
##### Params:
fn, string, File to be hashed.
##### Returns:
res, string or nil, Returns the hexadecimal checksum string on success, returns nil on failure.  
err, nil or string, Returns nil on success, returns error message string on failure.
### <font color="DeepSkyBlue">dates</font>
#### <font color="DeepSkyBlue">add</font>
##### Language: lua
##### Describe:
Get a date offset by a certain amount from a given date.
##### Params:
date, string, Date in yyyy-MM-dd format.  
year, number(int), Year offset.  
month, number(int), Month offset.  
day, number(int), Day offset.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">day</font>
##### Language: lua
##### Describe:
Get the day from a date.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, number(int), Day.
#### <font color="DeepSkyBlue">last_month</font>
##### Language: lua
##### Describe:
Get the same day of the previous month.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">last_quarter</font>
##### Language: lua
##### Describe:
Get the same day of the previous quarter.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">last_week</font>
##### Language: lua
##### Describe:
Get the same day of the previous week.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">last_year</font>
##### Language: lua
##### Describe:
Get the same day of the previous year.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">month</font>
##### Language: lua
##### Describe:
Get the month from a date.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, number(int), Month.
#### <font color="DeepSkyBlue">month_beg</font>
##### Language: lua
##### Describe:
Get the first day of the month.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">month_end</font>
##### Language: lua
##### Describe:
Get the last day of the month.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">new</font>
##### Language: lua
##### Describe:
Create a date.
##### Params:
year, number(int), Year.  
month, number(int), Month.  
day, number(int), Day.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">next_month</font>
##### Language: lua
##### Describe:
Get the same day of the next month.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">next_quarter</font>
##### Language: lua
##### Describe:
Get the same day of the next quarter.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">next_week</font>
##### Language: lua
##### Describe:
Get the same day of the next week.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">next_year</font>
##### Language: lua
##### Describe:
Get the same day of the next year.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">quarter</font>
##### Language: lua
##### Describe:
Get the quarter of a date.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, number(int), Quarter.
#### <font color="DeepSkyBlue">quarter_beg</font>
##### Language: lua
##### Describe:
Get the first day of the quarter.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">quarter_end</font>
##### Language: lua
##### Describe:
Get the last day of the quarter.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">sub</font>
##### Language: lua
##### Describe:
Calculate the number of days between two dates.
##### Params:
beg, string, Start date (yyyy-MM-dd).  
end, string, End date (yyyy-MM-dd).
##### Returns:
res, number(int), Number of days between.
#### <font color="DeepSkyBlue">today</font>
##### Language: lua
##### Describe:
Get today's date.
##### Params:
None.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">tomorrow</font>
##### Language: lua
##### Describe:
Get tomorrow's date.
##### Params:
None.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">week</font>
##### Language: lua
##### Describe:
Get the week number of a date.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, number(int), Week number.
#### <font color="DeepSkyBlue">week_beg</font>
##### Language: lua
##### Describe:
Get the first day of the week.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">week_end</font>
##### Language: lua
##### Describe:
Get the last day of the week.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">weekday</font>
##### Language: lua
##### Describe:
Get the day of the week for a date.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, number(int), Day of the week.
#### <font color="DeepSkyBlue">year</font>
##### Language: lua
##### Describe:
Get the year from a date.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, number(int), Year.
#### <font color="DeepSkyBlue">year_beg</font>
##### Language: lua
##### Describe:
Get the first day of the year.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">year_end</font>
##### Language: lua
##### Describe:
Get the last day of the year.
##### Params:
date, string, Date in yyyy-MM-dd format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">yesterday</font>
##### Language: lua
##### Describe:
Get yesterday's date.
##### Params:
None.
##### Returns:
res, string, Date in yyyy-MM-dd format.
### <font color="DeepSkyBlue">dbmsger</font>
#### <font color="DeepSkyBlue">client</font>
##### <font color="DeepSkyBlue">abstain</font>
###### Language: lua
###### Describe:
Release ownership of a message.  
After releasing ownership, the message can be consumed again.
###### Params:
id, number(int64), Message ID.  
version, number(int), Message version.  
pushTime, string, Message push time (yyyy-mm-dd HH:MM:SS).
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">ack</font>
###### Language: lua
###### Describe:
Acknowledge a message.
###### Params:
id, number(int64), Message ID.  
version, number(int), Message version.
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">pull</font>
###### Language: lua
###### Describe:
Pull messages.  
After pulling, the message ownership lasts for 5 minutes. If not acknowledged within 5 minutes, ownership will be revoked. To extend ownership, renew can be called within 5 minutes.
###### Params:
topics, table(1d), List of message topics.  
limit, number(int), Maximum number of messages to pull.
###### Returns:
res, table(1d) or nil, Call result data, returns the pulled message list on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">push</font>
###### Language: lua
###### Describe:
Push a message.  
The msg parameter is described as follows.  
msg.topic, string, message topic.  
msg.head, table(k/v), message header information, k/v are strings, can carry metadata like UA, optional.  
msg.body, table(k/v), message body.  
msg.relate_id, string, business ID associated with the message, used to ensure idempotency on the producer side, optional.  
msg.priority, number(int), priority ([0-99]), higher value means higher priority, optional.  
msg.push_time, string, message publish time (yyyy-mm-dd HH:MM:SS), used for delayed messages, optional.
###### Params:
msg, table(k/v), Message.
###### Returns:
res, number(int64), Call result data, returns message ID on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">renew</font>
###### Language: lua
###### Describe:
Renew message ownership.  
Renewing extends ownership to 5 minutes from the renewal time.
###### Params:
id, number(int64), Message ID.  
version, number(int), Message version.
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
#### <font color="DeepSkyBlue">is_duplicate_message</font>
##### Language: lua
##### Describe:
Check if the error is a duplicate message error.
##### Params:
err, string, Error description.
##### Returns:
res, bool, Result of the check.
##### Since: 0.1
#### <font color="DeepSkyBlue">is_no_longer_holding</font>
##### Language: lua
##### Describe:
Check if the error is a no-longer-holding error.
##### Params:
err, string, Error description.
##### Returns:
res, bool, Result of the check.
##### Since: 0.1
#### <font color="DeepSkyBlue">open</font>
##### Language: lua
##### Describe:
Open a DB connection.  
dbmsger supports connection reuse when calling open on an existing connection.  
The conf parameter is described as follows:  
conf.url, string, DB connection string.  
conf.table_name, string, message table name, default is dbmsger_msgs if not provided.
##### Params:
provider, string, DB type, supports mysql.  
conf, table(k/v), DB configuration.
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
### <font color="DeepSkyBlue">dingtalk</font>
#### <font color="DeepSkyBlue">send_action_card_msg</font>
##### Language: lua
##### Describe:
Send an action_card type message.  
Optional parameters are as follows:  
isAtAll, bool.  
atUserIds, table(1d).  
atMobiles, table(1d).  
btnOrientation, string, button layout (0 vertical, 1 horizontal).
##### Params:
accessToken, string, DingTalk robot access token.  
title, string, Display content shown in the first screen conversation.  
text, string, Markdown formatted message.  
singleTitle, string, Title of the single button.  
singleUrl, string, URL to jump to when the message is clicked.  
opts, table(k/v)(optional), Optional parameters (use as needed).
##### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">send_actions_card_msg</font>
##### Language: lua
##### Describe:
Send an actions_card type message.  
Optional parameters are as follows:  
isAtAll, bool.  
atUserIds, table(1d).  
atMobiles, table(1d).  
btnOrientation, string, button layout (0 vertical, 1 horizontal).
##### Params:
accessToken, string, DingTalk robot access token.  
title, string, Display content shown in the first screen conversation.  
text, string, Markdown formatted message.  
btns, table(1d), List of buttons (each element is {title="xxx", actionUrl="xxx"}).  
opts, table(k/v)(optional), Optional parameters (use as needed).
##### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">send_feed_card_msg</font>
##### Language: lua
##### Describe:
Send a feed_card type message.  
Optional parameters are as follows:  
isAtAll, bool.  
atUserIds, table(1d).  
atMobiles, table(1d).
##### Params:
accessToken, string, DingTalk robot access token.  
links, table(1d), List of links (each element is {title="xxx", messageUrl="xxx", picUrl="xxx"}).  
opts, table(k/v)(optional), Optional parameters (use as needed).
##### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">send_link_msg</font>
##### Language: lua
##### Describe:
Send a link type message.  
Optional parameters are as follows:  
isAtAll, bool.  
atUserIds, table(1d).  
atMobiles, table(1d).  
picUrl, string, image URL.
##### Params:
accessToken, string, DingTalk robot access token.  
title, string, Message title.  
text, string, Message content.  
messageUrl, string, URL to jump to when the message is clicked.  
opts, table(k/v)(optional), Optional parameters (use as needed).
##### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">send_markdown_msg</font>
##### Language: lua
##### Describe:
Send a markdown type message.  
Optional parameters are as follows:  
isAtAll, bool.  
atUserIds, table(1d).  
atMobiles, table(1d).
##### Params:
accessToken, string, DingTalk robot access token.  
title, string, Display content shown in the first screen conversation.  
text, string, Markdown formatted message.  
opts, table(k/v)(optional), Optional parameters (use as needed).
##### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">send_text_msg</font>
##### Language: lua
##### Describe:
Send a text type message.  
Optional parameters are as follows:  
isAtAll, bool.  
atUserIds, table(1d).  
atMobiles, table(1d).
##### Params:
accessToken, string, DingTalk robot access token.  
content, string, Message content.  
opts, table(k/v)(optional), Optional parameters (use as needed).
##### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.
##### Since: 0.1
### <font color="DeepSkyBlue">encodings</font>
#### <font color="DeepSkyBlue">base64_std_decode</font>
##### Language: lua
##### Describe:
Standard base64 decoding.
##### Params:
str, string, String to be decoded.
##### Returns:
res, string or nil, Returns the decoded string on success, returns nil on failure.  
err, nil or string, Returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">base64_std_encode</font>
##### Language: lua
##### Describe:
Standard base64 encoding.
##### Params:
str, string, String to be encoded.
##### Returns:
res, string, Encoded string.
### <font color="DeepSkyBlue">feishu</font>
#### <font color="DeepSkyBlue">client</font>
##### <font color="DeepSkyBlue">reply_msg</font>
###### Language: lua
###### Describe:
Reply to a message.
###### Params:
msgId, string, Message ID.  
msgType, string, Message type (text, post, image, file, audio, media, sticker, interactive, share_chat, share_user).  
content, string, Message content.
###### Returns:
msgId, string or nil, Returns the message ID on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">send_msg</font>
###### Language: lua
###### Describe:
Send a message.
###### Params:
receiveIdType, string, Type of recipient ID (open_id, union_id, user_id, email, chat_id).  
receiveId, string, Recipient ID.  
msgType, string, Message type (text, post, image, file, audio, media, sticker, interactive, share_chat, share_user, system).  
content, string, Message content.
###### Returns:
msgId, string or nil, Returns the message ID on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
#### <font color="DeepSkyBlue">open</font>
##### Language: lua
##### Describe:
Open a Feishu connection.
##### Params:
appId, string, Feishu appId.  
appSecret, string, Feishu appSecret.
##### Returns:
client, table, Client instance.
##### Since: 0.1
### <font color="DeepSkyBlue">filesystems</font>
#### <font color="DeepSkyBlue">append_file</font>
##### Language: lua
##### Describe:
Append content to a file, creates the file if it does not exist.
##### Params:
fn, string, File path.  
content, string, Content string to append.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">change_file_times</font>
##### Language: lua
##### Describe:
Change the timestamps of a file.
##### Params:
fn, string, File path.  
access, string, Access time in yyyy-MM-dd HH:mm:ss format.  
modify, string, Modification time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">copy_file</font>
##### Language: lua
##### Describe:
Copy a file, overwrites the destination file if it exists.
##### Params:
src, string, Source file path.  
dst, string, Destination file path.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">excel_2_csv</font>
##### Language: lua
##### Describe:
Convert Excel to CSV, supports .xlsx files, does not support .xls files.
##### Params:
src, string, Excel file path.  
sheet, string, Excel sheet name.  
dst, string, CSV file path.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">exist</font>
##### Language: lua
##### Describe:
Check if a file or folder path exists.
##### Params:
path, string, Path.
##### Returns:
res, bool, Returns true if exists, false otherwise.
#### <font color="DeepSkyBlue">join_path</font>
##### Language: lua
##### Describe:
Join path components.
##### Params:
dir, string, Parent directory path.  
name, string, File or folder name.
##### Returns:
res, string, Joined path.
#### <font color="DeepSkyBlue">list_dir</font>
##### Language: lua
##### Describe:
List directories under a given path, does not search subdirectories.
##### Params:
path, string, Folder path.
##### Returns:
ls, table(1d), List of folder names.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">list_file_with_exts</font>
##### Language: lua
##### Describe:
List files with specified extensions under a given path.
##### Params:
path, string, Folder path.  
sub, bool, Whether to search subdirectories.  
exts, string, File extensions, variable arguments, can specify multiple extensions, if not specified, list all files.
##### Returns:
ls, table(1d), List of file paths.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">list_file_with_pattern</font>
##### Language: lua
##### Describe:
List files matching a regex pattern under a given path.
##### Params:
path, string, Folder path.  
sub, bool, Whether to search subdirectories.  
pattern, string, Regex pattern to match file paths.
##### Returns:
ls, table(1d), List of file paths.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">mkdir</font>
##### Language: lua
##### Describe:
Create a directory path, supports creating multiple levels.
##### Params:
path, string, Folder path.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">read_excel</font>
##### Language: lua
##### Describe:
Read an Excel file, supports .xlsx files, does not support .xls files.
##### Params:
fn, string, File path.  
sheet, string, Sheet name.
##### Returns:
res, table(2d) or nil, Table data stored row by row, returns data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">read_excel_first_column</font>
##### Language: lua
##### Describe:
Read the first column of an Excel file, supports .xlsx files, does not support .xls files.
##### Params:
fn, string, File path.  
sheet, string, Sheet name.
##### Returns:
res, table(1d) or nil, Data from the first column, returns data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">read_file</font>
##### Language: lua
##### Describe:
Read a file.
##### Params:
fn, string, File path.
##### Returns:
res, string or nil, File content, returns content on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">read_file_info</font>
##### Language: lua
##### Describe:
Read file information.
##### Params:
fn, string, File path.
##### Returns:
res, table or nil, Returns file information on success, returns nil on failure.  
err, nil or string, Returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">read_json</font>
##### Language: lua
##### Describe:
Read a JSON file.
##### Params:
fn, string, File path.
##### Returns:
res, table or nil, Returns a table on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">remove</font>
##### Language: lua
##### Describe:
Delete a file or folder.
##### Params:
path, string, Path.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">rename</font>
##### Language: lua
##### Describe:
Rename a path.
##### Params:
old, string, Old path.  
new, string, New path.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">save_csv</font>
##### Language: lua
##### Describe:
Save data to a CSV file.
##### Params:
fn, string, File path.  
tb, table(2d), Table data stored row by row.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">save_excel</font>
##### Language: lua
##### Describe:
Save data to an Excel file, supports .xlsx files, does not support .xls files.
##### Params:
fn, string, File path.  
sheet, string, Sheet name.  
tb, table(2d), Table data stored row by row.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">save_file</font>
##### Language: lua
##### Describe:
Save content to a file, overwrites the file if it exists.
##### Params:
fn, string, File path.  
content, string, Content string to save.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">save_json</font>
##### Language: lua
##### Describe:
Save a table to a JSON file.
##### Params:
fn, string, File path.  
tb, table, Table object.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">search_file_line</font>
##### Language: lua
##### Describe:
Search for lines containing a specific string in files with a given extension under a path.
##### Params:
path, string, Folder path.  
sub, bool, Whether to search subdirectories.  
ext, string, File extension.  
content, string, Search string.
##### Returns:
ls, table(1d), List of matching file lines.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">split_path</font>
##### Language: lua
##### Describe:
Split a path into directory and name components.
##### Params:
path, string, Path.
##### Returns:
dir, string, Parent directory path.  
name, string, File name or last directory name.
### <font color="DeepSkyBlue">grpcaller</font>
#### <font color="DeepSkyBlue">client</font>
##### <font color="DeepSkyBlue">bench</font>
###### Language: lua
###### Describe:
Benchmark gRPC API.  
Each benchmark task contains the following information:  
api, string, API endpoint.  
meta, table(k/v), request metadata.  
req, table, request parameters.  
check_expr, string, expression to check business errors (e.g., code > 0).
###### Params:
parallel, number(int), Degree of parallelism.  
tps, number(float64), Requests per second.  
amount, number(int), Total number of calls.  
timeout, number(int), Timeout for each call in seconds.  
tasks, table(1d), List of benchmark tasks.
###### Returns:
res, table or nil, Returns benchmark result statistics on success, returns nil on failure.  
err, string, Error message, returns nil on success, returns error message on failure.
##### <font color="DeepSkyBlue">call</font>
###### Language: lua
###### Describe:
Call a gRPC API.
###### Params:
method, string, API method name.  
meta, table(k/v), Metadata, values are strings.  
req, string(json) or table(k/v), Request parameters.  
timeout, number(int), Call timeout in seconds.
###### Returns:
res, table(1d) or nil, Returns API response on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">close</font>
###### Language: lua
###### Describe:
Close the gRPC service connection.
###### Params:
None.
###### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">describe</font>
###### Language: lua
###### Describe:
Get description of a symbol (service, method, message).
###### Params:
name, string, Fully qualified symbol name.
###### Returns:
res, string or nil, Returns symbol description string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">input</font>
###### Language: lua
###### Describe:
Get the input type name of an API method.
###### Params:
method, string, API method name.
###### Returns:
res, string or nil, Returns the input type name on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">methods</font>
###### Language: lua
###### Describe:
Get the list of API methods for a service.
###### Params:
service, string, Service name.
###### Returns:
res, table(1d) or nil, Returns the list of method names on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">output</font>
###### Language: lua
###### Describe:
Get the output type name of an API method.
###### Params:
method, string, API method name.
###### Returns:
res, string or nil, Returns the output type name on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">services</font>
###### Language: lua
###### Describe:
Get the list of services.
###### Params:
None.
###### Returns:
res, table(1d) or nil, Returns the list of service names on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">template</font>
###### Language: lua
###### Describe:
Get a message template.
###### Params:
name, string, Fully qualified message symbol name.
###### Returns:
res, string or nil, Returns the message template string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
#### <font color="DeepSkyBlue">open</font>
##### Language: lua
##### Describe:
Open a gRPC service connection using server reflection (same as open_by_reflect).  
grpcaller supports connection reuse when opening an existing connection.
##### Params:
addr, string, Service address (ip:port).  
meta, table(k/v), Metadata, values are strings, used to control header-based routing for service discovery in cloud environments with multiple versions.  
opts, table(k/v)(optional), Optional connection parameters (use as needed).
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">open_by_proto</font>
##### Language: lua
##### Describe:
Open a gRPC service connection using proto files for service discovery.  
grpcaller supports connection reuse when opening an existing connection.
##### Params:
addr, string, Service address (ip:port).  
proto_dirs, table(1d), List of proto directory paths.  
proto_names, table(1d), List of proto file names.  
opts, table(k/v)(optional), Optional connection parameters (use as needed).
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">open_by_reflect</font>
##### Language: lua
##### Describe:
Open a gRPC service connection using server reflection for service discovery.  
grpcaller supports connection reuse when opening an existing connection.
##### Params:
addr, string, Service address (ip:port).  
meta, table(k/v), Metadata, values are strings, used to control header-based routing for service discovery in cloud environments with multiple versions.  
opts, table(k/v)(optional), Optional connection parameters (use as needed).
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
### <font color="DeepSkyBlue">https</font>
#### <font color="DeepSkyBlue">append</font>
##### Language: lua
##### Describe:
Append parameters to a URL.  
If the original URL contains a parameter with the same name, the original value will be replaced.
##### Params:
url, string, Original URL.  
params, table(1d), Parameters to append.
##### Returns:
res, string or nil, Returns the URL with appended parameters on success, returns nil on failure.  
err, string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">batch</font>
##### Language: lua
##### Describe:
Batch call HTTP APIs.
##### Params:
tasks, table, Batch requests, each request has the same parameters as the call function.  
parallel, number(int), Degree of parallelism (1-16).
##### Returns:
res, table or nil, Returns the corresponding result data for each request on success, returns nil on failure.  
err, string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">bench</font>
##### Language: lua
##### Describe:
Benchmark HTTP APIs.  
Each benchmark task contains the following information:  
method, string, HTTP method (GET/POST/...).  
url, string, API URL.  
header, table(k/v), additional HTTP headers.  
body, table(k/v) or table(1d) or string, request body data.  
check_expr, string, expression to check business errors (e.g., code > 0).
##### Params:
parallel, number(int), Degree of parallelism.  
tps, number(float64), Requests per second.  
amount, number(int), Total number of calls.  
timeout, number(int), Timeout for each call in seconds.  
tasks, table(1d), List of benchmark tasks.
##### Returns:
res, table or nil, Returns benchmark result statistics on success, returns nil on failure.  
err, string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">call</font>
##### Language: lua
##### Describe:
Call an HTTP API.
##### Params:
method, string, HTTP method (GET/POST).  
url, string, Request URL.  
header, table(k/v) or nil, Request headers, both keys and values must be strings.  
body, string or table or nil, Request body.  
timeout, number(int), Call timeout in seconds.  
pretty, bool, Whether to parse response data as a table, optional, default false (response as string), JSON responses can be parsed, non-JSON responses should not.
##### Returns:
res, string or table or nil, Returns response data as string or table on success, returns nil on failure.  
code, number(int), HTTP status code.  
err, string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">download</font>
##### Language: lua
##### Describe:
Download a file.
##### Params:
url, string, Request URL.  
header, table(k/v) or nil, Request headers, both keys and values must be strings.  
fn, string, Local file path to save the file.
##### Returns:
code, number(int), HTTP status code.  
err, string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">dump</font>
##### Language: lua
##### Describe:
Encode URL parameters.
##### Params:
params, table(k/v), Parameters to be encoded.
##### Returns:
res, string, Encoded parameter string.
#### <font color="DeepSkyBlue">parse</font>
##### Language: lua
##### Describe:
Parse URL parameters.
##### Params:
url, string, URL to be parsed.  
pretty, bool, Whether to parse JSON values into tables, optional, default false. if true, attempts to parse values enclosed by '{}', '[]', or equal to 'null' into tables.
##### Returns:
res, table(k, v) or nil, returns the parsed parameters table on success, returns nil on failure.  
err, string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">upload</font>
##### Language: lua
##### Describe:
Upload a file.
##### Params:
url, string, Request URL.  
header, table(k/v) or nil, Request headers, both keys and values must be strings.  
params, table(k, v) or nil, form parameters.  
fn, string, Local path of the file to upload.  
formname, string, Name of the file field in multipart/form-data.  
filename, string, Filename to be used in multipart/form-data.  
pretty, bool, Whether to parse response data as a table, optional, default false (response as string), JSON responses can be parsed, non-JSON responses should not.
##### Returns:
res, string or table or nil, Returns response data as string or table on success, returns nil on failure.  
code, number(int), HTTP status code.  
err, string, Error message, returns nil on success, returns error message on failure.
### <font color="DeepSkyBlue">jsons</font>
#### <font color="DeepSkyBlue">compare</font>
##### Language: lua
##### Describe:
Compare two JSON strings (experimental).
##### Params:
js1, string, First JSON string.  
js2, string, Second JSON string.
##### Returns:
res, table or nil, Returns a Lua table on success, returns nil on failure.  
err, string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">dump</font>
##### Language: lua
##### Describe:
Convert a table to JSON.
##### Params:
obj, table, Object to be serialized.  
pretty, bool, Whether to indent, optional, default false.
##### Returns:
res, string, JSON string.
#### <font color="DeepSkyBlue">get</font>
##### Language: lua
##### Describe:
Get a value from a JSON string using a path.
##### Params:
js, string, JSON string.  
path, string, Path.
##### Returns:
res, any, Retrieved value.
#### <font color="DeepSkyBlue">highlight</font>
##### Language: lua
##### Describe:
Highlight JSON (generates an HTML page).
##### Params:
titile, string, Title.  
content, string, JSON string.
##### Returns:
res, string, HTML page string.
#### <font color="DeepSkyBlue">mget</font>
##### Language: lua
##### Describe:
Get multiple values from a JSON string using a path mapping.
##### Params:
js, string, JSON string.  
projection, table(k/v), Path mapping.
##### Returns:
res, table, Retrieved table.
#### <font color="DeepSkyBlue">parse</font>
##### Language: lua
##### Describe:
Parse JSON to a table.
##### Params:
content, string, JSON string.  
pretty, bool, Whether to attempt parsing field values as JSON, optional, default false.
##### Returns:
res, table or nil, Returns a Lua table on success, returns nil on failure.  
err, string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">pretty</font>
##### Language: lua
##### Describe:
Pretty-print JSON (indent).
##### Params:
content, string, JSON string.
##### Returns:
res, string, Indented JSON string.
### <font color="DeepSkyBlue">kubernetes</font>
#### <font color="DeepSkyBlue">client</font>
##### <font color="DeepSkyBlue">call</font>
###### Language: lua
###### Describe:
Call a Kubernetes API.
###### Params:
method, string, HTTP method.  
api, string, API endpoint.  
body, string or table or nil, Request body.  
pretty, bool, Whether to parse response data as a table, optional, default false (response as string), JSON responses can be parsed, non-JSON responses should not.
###### Returns:
res, any or nil, Returns the response body on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">close</font>
###### Language: lua
###### Describe:
Close the Kubernetes API server connection.
###### Params:
None.
###### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">mcall</font>
###### Language: lua
###### Describe:
Call a Kubernetes API and return information based on a JSON projection.
###### Params:
method, string, HTTP method.  
api, string, API endpoint.  
body, string or table or nil, Request body.  
projection, table(k/v), Mapping for extracting and renaming fields from the response (key is JSON path, value is the mapped name), nil returns the full response.
###### Returns:
res, any or nil, Returns the response body on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">pcall</font>
###### Language: lua
###### Describe:
Call a Kubernetes API and return information based on a JSON path.
###### Params:
method, string, HTTP method.  
api, string, API endpoint.  
body, string or table or nil, Request body.  
path, string, JSON path to extract information, empty string returns the full response.
###### Returns:
res, any or nil, Returns the response body on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
#### <font color="DeepSkyBlue">open</font>
##### Language: lua
##### Describe:
Open a Kubernetes API server connection.  
kubernetes supports connection reuse when opening an existing connection.
##### Params:
kubeconfig, string, Kubernetes API server kubeconfig.
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">vendor</font>
##### Language: lua
##### Describe:
Cloud vendor (unknown/alibaba/bytedance/tencent/huawei/amazon/microsoft/google).
##### Params:
None.
##### Returns:
res, string, Cloud vendor.
##### Since: 0.1
### <font color="DeepSkyBlue">logger</font>
#### <font color="DeepSkyBlue">debug</font>
##### Language: lua
##### Describe:
Print debug level message.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
str, string, String information.
##### Returns:
None.
##### Since: 0.1
#### <font color="DeepSkyBlue">debugd</font>
##### Language: lua
##### Describe:
Extended print debug level message.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
detail, table(k/v), Additional extra information.
##### Returns:
None.
##### Since: 0.1
#### <font color="DeepSkyBlue">debugf</font>
##### Language: lua
##### Describe:
Formatted print debug level message.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
pattern, string, String format pattern.  
args, string(variable arguments), Arguments to be formatted (multiple arguments can be specified).
##### Returns:
None.
##### Since: 0.1
#### <font color="DeepSkyBlue">error</font>
##### Language: lua
##### Describe:
Print error level message.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
str, string, String information.
##### Returns:
None.
##### Since: 0.1
#### <font color="DeepSkyBlue">errord</font>
##### Language: lua
##### Describe:
Extended print error level message.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
detail, table(k/v), Additional extra information.
##### Returns:
None.
##### Since: 0.1
#### <font color="DeepSkyBlue">errorf</font>
##### Language: lua
##### Describe:
Formatted print error level message.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
pattern, string, String format pattern.  
args, string(variable arguments), Arguments to be formatted (multiple arguments can be specified).
##### Returns:
None.
##### Since: 0.1
#### <font color="DeepSkyBlue">info</font>
##### Language: lua
##### Describe:
Print info level message.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
str, string, String information.
##### Returns:
None.
##### Since: 0.1
#### <font color="DeepSkyBlue">infod</font>
##### Language: lua
##### Describe:
Extended print info level message.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
detail, table(k/v), Additional extra information.
##### Returns:
None.
##### Since: 0.1
#### <font color="DeepSkyBlue">infof</font>
##### Language: lua
##### Describe:
Formatted print info level message.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
pattern, string, String format pattern.  
args, string(variable arguments), Arguments to be formatted (multiple arguments can be specified).
##### Returns:
None.
##### Since: 0.1
#### <font color="DeepSkyBlue">raw</font>
##### Language: lua
##### Describe:
Print raw message directly.
##### Params:
pusher, string, Message publisher.  
topics, string, Message topic (multiple topics can be specified, separated by '|').  
msg, table(k/v), Message.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">trace</font>
##### Language: lua
##### Describe:
Print call stack information.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
detail, table(k/v), Additional extra information.  
level, number(int), Call stack level, level >= 1, optional, default 1.
##### Returns:
None.
##### Since: 0.1
#### <font color="DeepSkyBlue">warn</font>
##### Language: lua
##### Describe:
Print warn level message.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
str, string, String information.
##### Returns:
None.
##### Since: 0.1
#### <font color="DeepSkyBlue">warnd</font>
##### Language: lua
##### Describe:
Extended print warn level message.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
detail, table(k/v), Additional extra information.
##### Returns:
None.
##### Since: 0.1
#### <font color="DeepSkyBlue">warnf</font>
##### Language: lua
##### Describe:
Formatted print warn level message.
##### Params:
topics, string, Message topic (multiple topics can be specified, separated by '|').  
keys, string, Message keywords (multiple keywords can be specified, separated by '|').  
pattern, string, String format pattern.  
args, string(variable arguments), Arguments to be formatted (multiple arguments can be specified).
##### Returns:
None.
##### Since: 0.1
### <font color="DeepSkyBlue">mails</font>
#### <font color="DeepSkyBlue">send</font>
##### Language: lua
##### Describe:
Send an email.
##### Params:
user, string, Sender email address.  
password, string, Sender email password.  
host, string, Sender mail server hostname.  
port, number(int), Sender mail server port.  
to, lua table(1d), List of recipient email addresses (at least one).  
cc, lua table(1d), List of CC email addresses (empty table if none).  
bcc, lua table(1d), List of BCC email addresses (empty table if none).  
subject, string, Email subject.  
body, string, Email body.  
content_type, string, Email content type.  
attachment_with_file, bool, Whether there is an attachment.  
attachment_name, string, Attachment file name.  
attachment_content_type, string, Attachment content type.  
attachment_path, string, Attachment file path.
##### Returns:
err, nil or string, Send error message, returns nil on success, returns error message string on failure.
### <font color="DeepSkyBlue">mongodb</font>
#### <font color="DeepSkyBlue">binary</font>
##### Language: lua
##### Describe:
Create a binary extended type.  
Currently, UUID is parsed as a hyphen-separated string format, all other types are parsed as base64 strings.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
subtype, string, Subtype (generic, function, binary_old, uuid_old, uuid, md5, encrypted, column, user_defined).  
data, string, String data.
##### Returns:
res, binary, Result value.
#### <font color="DeepSkyBlue">bson_d</font>
##### Language: lua
##### Describe:
Create a bson_d extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
bsones, bson_e, bson_e parameters, variable arguments, can specify multiple as needed.
##### Returns:
res, bson_d, Result value.
#### <font color="DeepSkyBlue">bson_e</font>
##### Language: lua
##### Describe:
Create a bson_e extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
key, string, bson.E.Key.  
value, any, bson.E.Value.
##### Returns:
res, bson_e, Result value.
#### <font color="DeepSkyBlue">client</font>
##### <font color="DeepSkyBlue">database</font>
###### Language: lua
###### Describe:
Get a database.
###### Params:
database, string, Database name.
###### Returns:
db, table, Database instance.
###### Since: 0.1
###### <font color="DeepSkyBlue">aggregate</font>
###### Language: lua
###### Describe:
Aggregate data in the database.  
Optional parameters are as follows:  
let, any.  
hint, any.  
comment, string.  
collation, table(kv).  
maxTime, number(int).  
maxAwaitTime, number(int).  
batchSize, number(int32).  
allowDiskUse, bool.  
bypassDocumentValidation, bool.  
custom, table(k/v).
###### Params:
pipeline, table, Pipeline.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table(1d) or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">collection</font>
###### Language: lua
###### Describe:
Get a collection.
###### Params:
collection, string, Collection name.  
###### Returns:
coll, table, Collection instance.  
###### Since: 0.1
###### <font color="DeepSkyBlue">aggregate</font>
###### Language: lua
###### Describe:
Aggregate data in the collection.  
Optional parameters are as follows:  
let, any.  
hint, any.  
comment, string.  
collation, table(kv).  
maxTime, number(int).  
maxAwaitTime, number(int).  
batchSize, number(int32).  
allowDiskUse, bool.  
bypassDocumentValidation, bool.  
custom, table(k/v).
###### Params:
pipeline, table, Pipeline.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table(1d) or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">count_documents</font>
###### Language: lua
###### Describe:
Count documents.  
Optional parameters are as follows:  
hint, any.  
comment, string.  
collation, table(k/v).  
maxTime, number(int), in seconds.  
limit, number(int64).  
skip, number(int64).
###### Params:
filter, table(k, v), filter.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, number(int64) or nil, Returns the document count on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">delete_many</font>
###### Language: lua
###### Describe:
Delete multiple documents.  
Optional parameters are as follows:  
let, any.  
comment, string.  
hint, any.  
collation, table(k/v).
###### Params:
filter, table(k/v), Filter.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, number(int64) or nil, Returns the number of deleted documents on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">delete_one</font>
###### Language: lua
###### Describe:
Delete one document.  
Optional parameters are as follows:  
let, any.  
comment, string.  
hint, any.  
collation, table(k/v).
###### Params:
filter, table(k/v), Filter.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, number(int64) or nil, Returns the number of deleted documents on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">distinct</font>
###### Language: lua
###### Describe:
Get distinct values of a specified field in the collection.  
Optional parameters are as follows:  
comment, string.  
collation, table(k/v).  
maxTime, number(int).
###### Params:
field, string, Field to deduplicate.  
filter, table(k, v), filter.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table(1d) or nil, Returns the list of distinct values on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">drop</font>
###### Language: lua
###### Describe:
Drop the collection.
###### Params:
None.
###### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">estimated_document_count</font>
###### Language: lua
###### Describe:
Estimate document count.  
Optional parameters are as follows:  
comment, string.  
maxTime, number(int), in seconds.
###### Params:
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, number(int64) or nil, Returns the estimated document count on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">find</font>
###### Language: lua
###### Describe:
Find documents.  
Optional parameters are as follows:  
projection, table(k/v).  
skip, number(int64).  
limit, number(int64).  
sort, table(k/v).  
collation, table(kv).  
allowDiskUse, bool.  
allowPartialResults, bool.  
batchSize, number(int32).  
comment, string.  
maxTime, number(int).  
maxAwaitTime, number(int).  
min, any.  
max, any.  
hint, any.  
let, any.  
returnKey, bool.  
showRecordID, bool.  
noCursorTimeout, bool.  
cursorType, number(int8).
###### Params:
filter, table(k/v), Filter.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table(1d) or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">find_and_dump</font>
###### Language: lua
###### Describe:
Find documents and return results as JSON, parameters are the same as find.  
Optional parameters are as follows:  
projection, table(k/v).  
skip, number(int64).  
limit, number(int64).  
sort, table(k/v).  
collation, table(kv).  
allowDiskUse, bool.  
allowPartialResults, bool.  
batchSize, number(int32).  
comment, string.  
maxTime, number(int).  
maxAwaitTime, number(int).  
min, any.  
max, any.  
hint, any.  
let, any.  
returnKey, bool.  
showRecordID, bool.  
noCursorTimeout, bool.  
cursorType, number(int8).
###### Params:
filter, table(k/v), Filter.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, string or nil, Returns result data as JSON string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">find_for_each</font>
###### Language: lua
###### Describe:
Find documents and iterate over each document using a callback.  
Optional parameters are as follows:  
projection, table(k/v).  
skip, number(int64).  
limit, number(int64).  
sort, table(k/v).  
collation, table(kv).  
allowDiskUse, bool.  
allowPartialResults, bool.  
batchSize, number(int32).  
comment, string.  
maxTime, number(int).  
maxAwaitTime, number(int).  
min, any.  
max, any.  
hint, any.  
let, any.  
returnKey, bool.  
showRecordID, bool.  
noCursorTimeout, bool.  
cursorType, number(int8).
###### Params:
filter, table(k/v), Filter.  
callback, func(doc string), Callback function.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">find_one</font>
###### Language: lua
###### Describe:
Find one document.  
Optional parameters are as follows:  
projection, table(k/v).  
skip, number(int64).  
sort, table(k/v).  
collation, table(kv).  
allowPartialResults, bool.  
comment, string.  
maxTime, number(int).  
min, any.  
max, any.  
hint, any.  
returnKey, bool.  
showRecordID, bool.
###### Params:
filter, table(k/v), Filter.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">find_one_and_delete</font>
###### Language: lua
###### Describe:
Find and delete one document.  
Optional parameters are as follows:  
projection, table(k/v).  
sort, table(k/v).  
collation, table(kv).  
comment, string.  
maxTime, number(int).  
hint, any.
###### Params:
filter, table(k/v), Filter.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">find_one_and_dump</font>
###### Language: lua
###### Describe:
Find one document and return result as JSON, parameters are the same as find_one.  
Optional parameters are as follows:  
projection, table(k/v).  
skip, number(int64).  
sort, table(k/v).  
collation, table(kv).  
allowPartialResults, bool.  
comment, string.  
maxTime, number(int).  
min, any.  
max, any.  
hint, any.  
returnKey, bool.  
showRecordID, bool.
###### Params:
filter, table(k/v), Filter.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, string or nil, Returns result data as JSON string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">find_one_and_replace</font>
###### Language: lua
###### Describe:
Find and replace one document.  
Optional parameters are as follows:  
projection, table(k/v).  
sort, table(k/v).  
collation, table(kv).  
comment, string.  
maxTime, number(int).  
hint, any.  
let, any.  
bypassDocumentValidation, bool.  
returnDocument, number(int8).  
upsert, bool.
###### Params:
filter, table(k/v), Filter.  
replace, table(k/v), Replacement document.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">find_one_and_update</font>
###### Language: lua
###### Describe:
Find and update one document.  
Optional parameters are as follows:  
projection, table(k/v).  
sort, table(k/v).  
collation, table(kv).  
comment, string.  
maxTime, number(int).  
hint, any.  
let, any.  
bypassDocumentValidation, bool.  
returnDocument, number(int8).  
upsert, bool.
###### Params:
filter, table(k/v), Filter.  
update, table(k/v), Update data.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">indexes</font>
###### Language: lua
###### Describe:
Get indexes.
###### Params:
None.
###### Returns:
res, table, Index object.  
###### Since: 0.1
###### <font color="DeepSkyBlue">list</font>
###### Language: lua
###### Describe:
List indexes in the collection.
###### Params:
None.
###### Returns:
res, table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">insert_many</font>
###### Language: lua
###### Describe:
Insert multiple documents.  
Optional parameters are as follows:  
comment, string.  
bypassDocumentValidation, bool.  
ordered, bool.
###### Params:
docs, table(1d), Documents.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table(1d) or nil, Returns the list of inserted document IDs on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">insert_one</font>
###### Language: lua
###### Describe:
Insert one document.  
Optional parameters are as follows:  
comment, string.  
bypassDocumentValidation, bool.
###### Params:
doc, table, Document.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, any or nil, Returns the inserted document ID on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">replace_one</font>
###### Language: lua
###### Describe:
Replace one document.  
Optional parameters are as follows:  
let, any.  
comment, string.  
bypassDocumentValidation, bool.  
hint, any.  
upsert, bool.  
collation, table(k/v).
###### Params:
filter, table(k/v), Filter.  
replace, table(k, v), replacement document.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table(k/v) or nil, Returns operation result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">schema</font>
###### Language: lua
###### Describe:
Get the schema of the collection.
###### Params:
sample, number(int), Sample size, optional, default 500.  
timeout, number(int), Timeout in seconds, optional, default 10s.  
###### Returns:
res, table(k/v) or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">update_by_id</font>
###### Language: lua
###### Describe:
Update a document by ID.  
Optional parameters are as follows:  
let, any.  
comment, string.  
bypassDocumentValidation, bool.  
hint, any.  
upsert, bool.  
collation, table(k/v).
###### Params:
id, any, _id.  
update, table(k/v), Update data.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table(k/v) or nil, Returns operation result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">update_many</font>
###### Language: lua
###### Describe:
Update multiple documents.  
Optional parameters are as follows:  
let, any.  
comment, string.  
bypassDocumentValidation, bool.  
hint, any.  
upsert, bool.  
collation, table(k/v).
###### Params:
filter, table(k/v), Filter.  
update, table(k/v), Update data.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table(k/v) or nil, Returns operation result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">update_one</font>
###### Language: lua
###### Describe:
Update one document.  
Optional parameters are as follows:  
let, any.  
comment, string.  
bypassDocumentValidation, bool.  
hint, any.  
upsert, bool.  
collation, table(k/v).
###### Params:
filter, table(k/v), Filter.  
update, table(k/v), Update data.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table(k/v) or nil, Returns operation result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">list_collection_names</font>
###### Language: lua
###### Describe:
List collection names in the database.  
Optional parameters are as follows:  
batchSize, number(int32).  
authorizedCollections, bool.  
nameOnly, bool.  
Note: Currently the filter parameter only works with types.mongo_bsond(), needs improvement.
###### Params:
filter, table(k/v), Filter.  
opts, table(k/v)(optional), Optional parameters (use as needed).  
###### Returns:
res, table(1d) or nil, Returns the list of collection names on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">run_command</font>
###### Language: lua
###### Describe:
Execute a MongoDB command.
###### Params:
command, table(k/v), Command.  
###### Returns:
res, table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
##### <font color="DeepSkyBlue">list_database_names</font>
###### Language: lua
###### Describe:
List database names.  
Optional parameters are as follows:  
authorizedDatabases, bool.  
nameOnly, bool.  
Note: Currently the filter parameter only works with types.mongo_bsond(), needs improvement.
###### Params:
filter, table(k/v), Filter.  
opts, table(k/v)(optional), Optional parameters (use as needed).
###### Returns:
res, table(1d) or nil, Returns the list of database names on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">ping</font>
###### Language: lua
###### Describe:
Ping the MongoDB server.
###### Params:
None.
###### Returns:
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
#### <font color="DeepSkyBlue">date_time</font>
##### Language: lua
##### Describe:
Create a date_time extended type, parses yyyy-MM-dd HH:mm:ss format time to MongoDB datetime.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, date_time, Result value.
#### <font color="DeepSkyBlue">decimal_128</font>
##### Language: lua
##### Describe:
Create a decimal_128 extended type from a number or string.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number or string, Numeric value.
##### Returns:
res, decimal_128, Result value.
#### <font color="DeepSkyBlue">is_err_nodocument</font>
##### Language: lua
##### Describe:
Check if the error is a no document found error.
##### Params:
err, string, Error description string.
##### Returns:
res, bool, Result of the check.
##### Since: 0.1
#### <font color="DeepSkyBlue">number_long</font>
##### Language: lua
##### Describe:
Create a number_long extended type from a numeric string.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, string, Numeric string.
##### Returns:
res, number_long, Result value.
#### <font color="DeepSkyBlue">object_id</font>
##### Language: lua
##### Describe:
Create an object_id extended type from a hex string ID.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, string, Hex string ID.
##### Returns:
res, object_id, Result value.
#### <font color="DeepSkyBlue">open</font>
##### Language: lua
##### Describe:
Open a MongoDB connection.  
mongodb supports connection reuse when calling open on an existing connection.  
Optional connection parameters are as follows:  
auth, table(k/v).  
appName, string.  
compressors, table(1d).  
direct, bool.  
connectTimeout, number(int), connection timeout in seconds.  
disableOCSPEndpointCheck, bool.  
heartbeatInterval, int, heartbeat interval in seconds.  
hosts, table(1d), server host list.  
loadBalanced, bool.  
localThreshold, int, in seconds.  
maxConnecting, number(uint64).  
maxConnIdleTime, number(int), in seconds.  
minPoolSize, number(uint64).  
maxPoolSize, number(uint64).  
readConcern, table(k/v).  
readPreference, string, options: primary, primaryPreferred, secondary, secondaryPreferred, nearest, default primary.  
writeConcern, table(k/v).  
replicaSet, string.  
retryReads, bool.  
retryWrites, bool.  
serverSelectionTimeout, number(int), in seconds.  
socketTimeout, number(int), in seconds.  
SRVMaxHosts, int.  
SRVServiceName, string.  
timeout, number(int), in seconds.  
zlibLevel, int.  
zstdLevel, int.
##### Params:
uri, string, Connection string.  
opts, table(k/v)(optional), Optional connection parameters (use as needed).
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
### <font color="DeepSkyBlue">nets</font>
#### <font color="DeepSkyBlue">ip_geo_district</font>
##### Language: lua
##### Describe:
IP geolocation.
##### Params:
ip, string, IP address.
##### Returns:
res, table or nil, Returns geolocation result on success, returns nil on failure.  
err, string, Error message, returns nil on success, returns error message on failure.
### <font color="DeepSkyBlue">pbparser</font>
#### <font color="DeepSkyBlue">new</font>
##### Language: lua
##### Describe:
Create a new protobuf parser.
##### Params:
pdesc, string, Proto descriptor file.
##### Returns:
parser, table or nil, Returns the parser table on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">parser</font>
##### <font color="DeepSkyBlue">convert_bin_to_json</font>
###### Language: lua
###### Describe:
Convert binary protobuf data to JSON.
###### Params:
content, string, Data to convert.  
messageTypeName, string, Message type name.
###### Returns:
res, string or nil, Returns the JSON string on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">convert_bin_to_text</font>
###### Language: lua
###### Describe:
Convert binary protobuf data to text format.
###### Params:
content, string, Data to convert.  
messageTypeName, string, Message type name.
###### Returns:
res, string or nil, Returns the text string on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">convert_json_to_bin</font>
###### Language: lua
###### Describe:
Convert JSON protobuf data to binary format.
###### Params:
content, string, Data to convert.  
messageTypeName, string, Message type name.
###### Returns:
res, string or nil, Returns the binary string on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">convert_json_to_text</font>
###### Language: lua
###### Describe:
Convert JSON protobuf data to text format.
###### Params:
content, string, Data to convert.  
messageTypeName, string, Message type name.
###### Returns:
res, string or nil, Returns the text string on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">convert_text_to_bin</font>
###### Language: lua
###### Describe:
Convert text protobuf data to binary format.
###### Params:
content, string, Data to convert.  
messageTypeName, string, Message type name.
###### Returns:
res, string or nil, Returns the binary string on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">convert_text_to_json</font>
###### Language: lua
###### Describe:
Convert text protobuf data to JSON format.
###### Params:
content, string, Data to convert.  
messageTypeName, string, Message type name.
###### Returns:
res, string or nil, Returns the JSON string on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">get_message_schema</font>
###### Language: lua
###### Describe:
Get message schema.
###### Params:
messageTypeName, string, Message type name.
###### Returns:
res, table or nil, Returns the schema table on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">list_message_types</font>
###### Language: lua
###### Describe:
List message types.
###### Params:
None.
###### Returns:
res, table(1d), List of message types.
##### <font color="DeepSkyBlue">unmarshal_bin</font>
###### Language: lua
###### Describe:
Parse protobuf binary data.
###### Params:
content, string, Data to parse.  
messageTypeName, string, Message type name.
###### Returns:
res, table or nil, Returns the parsed table on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">unmarshal_json</font>
###### Language: lua
###### Describe:
Parse protobuf JSON data.
###### Params:
content, string, Data to parse.  
messageTypeName, string, Message type name.
###### Returns:
res, table or nil, Returns the parsed table on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">unmarshal_text</font>
###### Language: lua
###### Describe:
Parse protobuf text data.
###### Params:
content, string, Data to parse.  
messageTypeName, string, Message type name.
###### Returns:
res, table or nil, Returns the parsed table on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
### <font color="DeepSkyBlue">pulsar</font>
#### <font color="DeepSkyBlue">client</font>
##### <font color="DeepSkyBlue">consumer</font>
###### <font color="DeepSkyBlue">consume</font>
###### Language: lua
###### Describe:
Consume messages.
###### Params:
process, function(msg) bool, Message processing function (msg is a table {topic=string, producer_name=string, properties=table(k/v), payload=string, id=string, publish_time=int(unix ms), redelivery_count=int}).  
###### Returns:
err, string or nil, Error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">create_consumer</font>
###### Language: lua
###### Describe:
Create a Pulsar consumer.
###### Params:
name, string, Name.  
type, string, Type (exclusive, shared, failover, key_shared).  
topics, table(1d), Message topics (can subscribe to multiple).
###### Returns:
consumer, table, Consumer instance.  
err, string or nil, Error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">create_producer</font>
###### Language: lua
###### Describe:
Create a Pulsar producer.
###### Params:
name, string, Name.  
topic, string, Message topic.
###### Returns:
producer, table, Producer instance.  
err, string or nil, Error message, returns nil on success, returns error message string on failure.
##### <font color="DeepSkyBlue">producer</font>
###### <font color="DeepSkyBlue">publish</font>
###### Language: lua
###### Describe:
Publish a message.
###### Params:
properties, table(k/v), Message properties (both keys and values are strings).  
payload, string, Message body.  
deliverAfter, number(int), Delivery delay in ms, 0 for no delay.  
deliverAt, string, Delivery time (yyyy-mm-dd HH:MM:SS), empty for immediate delivery.  
###### Returns:
messageId, string or nil, Message ID, returns message ID on success, returns nil on failure.  
err, string or nil, Error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">open</font>
##### Language: lua
##### Describe:
Create a Pulsar client.
##### Params:
url, string, Pulsar connection URL.  
token, string, Access token.
##### Returns:
client, table or nil, Client instance, returns client instance on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
### <font color="DeepSkyBlue">rands</font>
#### <font color="DeepSkyBlue">hexstr</font>
##### Language: lua
##### Describe:
Generate a random hex string.
##### Params:
n, number(int), Number of bytes before hex encoding.
##### Returns:
res, string, Hex string.
#### <font color="DeepSkyBlue">shuffle</font>
##### Language: lua
##### Describe:
Shuffle a table.
##### Params:
t, table(1d), Table to be shuffled.
##### Returns:
None.
#### <font color="DeepSkyBlue">uuid</font>
##### Language: lua
##### Describe:
Generate a UUID.
##### Params:
None.
##### Returns:
res, string, UUID string.
### <font color="DeepSkyBlue">redis</font>
#### <font color="DeepSkyBlue">client</font>
##### <font color="DeepSkyBlue">bgsave</font>
###### Language: lua
###### Describe:
Asynchronously persist current cache data to disk in the background.
###### Params:
None.
###### Returns:
res, string or nil, Returns OK on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">bloom</font>
###### Language: lua
###### Describe:
Get a Bloom filter.
###### Params:
n, number(int), Expected number of elements.  
p, number(float), False positive rate.  
b, number(int), Number of bytes per counter (0-4, 0 uses bits, 1-4 uses strings).
###### Returns:
res, table, Bloom filter instance.
###### Since: 0.1
###### <font color="DeepSkyBlue">add</font>
###### Language: lua
###### Describe:
Add an element to the Bloom filter.
###### Params:
key, string, Key.  
data, string, Element data.  
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">del</font>
###### Language: lua
###### Describe:
Delete an element from the Bloom filter.
###### Params:
key, string, Key.  
data, string, Element data.  
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">has</font>
###### Language: lua
###### Describe:
Check if an element may exist in the Bloom filter.
###### Params:
key, string, Key.  
data, string, Element data.  
###### Returns:
res, bool or nil, Returns true or false on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">reset</font>
###### Language: lua
###### Describe:
Reset the Bloom filter.
###### Params:
key, string, Key.  
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
##### <font color="DeepSkyBlue">blpop</font>
###### Language: lua
###### Describe:
Blocking left pop from multiple list keys.
###### Params:
timeout, number(int), Blocking timeout in seconds.  
keys, string(variable arguments), Cache keys (multiple can be provided).
###### Returns:
res, table(1d) or nil, Returns {key, value} on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">brpop</font>
###### Language: lua
###### Describe:
Blocking right pop from multiple list keys.
###### Params:
timeout, number(int), Blocking timeout in seconds.  
keys, string(variable arguments), Cache keys (multiple can be provided).
###### Returns:
res, table(1d) or nil, Returns {key, value} on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">del</font>
###### Language: lua
###### Describe:
Delete multiple keys.
###### Params:
keys, string, Cache keys, variable arguments, can specify multiple.
###### Returns:
res, number(int64), Returns the number of deleted keys on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">exists</font>
###### Language: lua
###### Describe:
Check if keys exist.
###### Params:
keys, string(variable arguments), Keys to check (multiple can be provided).
###### Returns:
num, number or nil, Returns the number of existing keys on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">expire</font>
###### Language: lua
###### Describe:
Set TTL for a key.
###### Params:
key, string, Cache key.  
ttl, number(int), TTL in seconds, <= 0 means persistent.
###### Returns:
res, bool, Returns true on success, false on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get</font>
###### Language: lua
###### Describe:
Get a string value.
###### Params:
key, string, Cache key.
###### Returns:
res, string or nil, Returns the cached data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">getbit</font>
###### Language: lua
###### Describe:
Get the bit value at a specific offset in a string key.
###### Params:
key, string, Cache key.  
offset, number(int64), Offset.
###### Returns:
res, number(int64) or nil, Returns the bit value at the specified offset on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">hget</font>
###### Language: lua
###### Describe:
Get a field from a hash.
###### Params:
key, string, Cache key.  
field, string, Field name in the hash.
###### Returns:
res, string or nil, Returns the value from the hash on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">hgetall</font>
###### Language: lua
###### Describe:
Get all fields and values from a hash.
###### Params:
key, string, Cache key.
###### Returns:
res, table(k/v) or nil, Returns all fields and values from the hash on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">hmget</font>
###### Language: lua
###### Describe:
Get multiple fields from a hash.
###### Params:
key, string, Cache key.  
fields, string(variable arguments), Field names in the hash (multiple can be provided).
###### Returns:
res, table(1d) or nil, Returns a list of field-value pairs {field, value, field, value, ...} on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">hmset</font>
###### Language: lua
###### Describe:
Set multiple fields in a hash.
###### Params:
key, string, Cache key.  
fvs, any, Variable arguments.
###### Returns:
res, bool, Returns true on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">hset</font>
###### Language: lua
###### Describe:
Set a field in a hash.
###### Params:
key, string, Cache key.  
field, string, Field name in the hash.  
value, any, Value.
###### Returns:
res, number(int64), Returns true on success, false on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">incr</font>
###### Language: lua
###### Describe:
Increment a string value that represents a number by 1.
###### Params:
key, string, Cache key.
###### Returns:
res, number(int64), Returns the incremented value on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">limiter</font>
###### Language: lua
###### Describe:
Get a rate limiter.
###### Params:
None.
###### Returns:
res, table, Rate limiter instance.
###### Since: 0.1
###### <font color="DeepSkyBlue">allow</font>
###### Language: lua
###### Describe:
Rate limit update and check (increment by 1).
###### Params:
key, string, Key.  
rate, string, Rate (e.g., 10/s, 30/m, 360/h).  
###### Returns:
res, table or nil, Current state, returns table on success (res.Remaining <= 0 triggers rate limit), returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">allow_at_most</font>
###### Language: lua
###### Describe:
Rate limit update and check (increment by at most n).
###### Params:
key, string, Key.  
rate, string, Rate (e.g., 10/s, 30/m, 360/h).  
n, number(int), Maximum increment.  
###### Returns:
res, table or nil, Current state, returns table on success (res.Remaining <= 0 triggers rate limit), returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">allow_n</font>
###### Language: lua
###### Describe:
Rate limit update and check (increment by n).
###### Params:
key, string, Key.  
rate, string, Rate (e.g., 10/s, 30/m, 360/h).  
n, number(int), Increment.  
###### Returns:
res, table or nil, Current state, returns table on success (res.Remaining <= 0 triggers rate limit), returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
###### <font color="DeepSkyBlue">reset</font>
###### Language: lua
###### Describe:
Reset the rate limiter.
###### Params:
key, string, Key.  
###### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.  
###### Since: 0.1
##### <font color="DeepSkyBlue">llen</font>
###### Language: lua
###### Describe:
Get the length of a list.
###### Params:
key, string, Cache key.
###### Returns:
res, number(int64) or nil, Returns the length of the list on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">lpop</font>
###### Language: lua
###### Describe:
Left pop an element from a list.
###### Params:
key, string, Cache key.
###### Returns:
res, string or nil, Returns the popped element on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">lpush</font>
###### Language: lua
###### Describe:
Left push elements to a list.
###### Params:
key, string, Cache key.  
values, any(variable arguments), Data to push (multiple can be provided).
###### Returns:
res, number(int64) or nil, Returns the new length of the list on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">lrange</font>
###### Language: lua
###### Describe:
Get a range of elements from a list.
###### Params:
key, string, Cache key.  
start, number(int64), Start index, 0 is the first element.  
end, number(int64), End index, inclusive, -1 is the last element.
###### Returns:
res, table(1d) or nil, Returns the list of elements on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">publish</font>
###### Language: lua
###### Describe:
Publish a message to a channel.
###### Params:
channel, string, Channel.  
message, any, Message.
###### Returns:
res, number(int64) or nil, Returns the number of subscribers that received the message on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">rpop</font>
###### Language: lua
###### Describe:
Right pop an element from a list.
###### Params:
key, string, Cache key.
###### Returns:
res, string or nil, Returns the popped element on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">rpush</font>
###### Language: lua
###### Describe:
Right push elements to a list.
###### Params:
key, string, Cache key.  
values, any(variable arguments), Data to push (multiple can be provided).
###### Returns:
res, number(int64) or nil, Returns the new length of the list on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">sadd</font>
###### Language: lua
###### Describe:
Add one or more elements to a set.
###### Params:
key, string, Cache key.  
members, string(variable arguments), Elements to add.
###### Returns:
res, number(int64), Returns the number of new elements added on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">scan</font>
###### Language: lua
###### Describe:
Scan keys in the cache.
###### Params:
cursor, number(int64), Scan cursor.  
match, string, Key match pattern.  
count, number(int64), Expected number of keys to return.
###### Returns:
keys, table(1d) or nil, Returns the list of scanned keys on success, returns nil on failure.  
cursor, number(int64), Returns the cursor for the next scan on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">scard</font>
###### Language: lua
###### Describe:
Get the cardinality of a set.
###### Params:
key, string, Cache key.
###### Returns:
res, number(int64), Returns the number of elements in the set on success (0 if key does not exist), returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">sdiff</font>
###### Language: lua
###### Describe:
Get the difference between the first set and other sets.
###### Params:
keys, string(variable arguments), Cache keys.
###### Returns:
res, table(1d), Returns the list of elements in the difference on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">sdiffstore</font>
###### Language: lua
###### Describe:
Compute the difference between the first set and other sets and store in the destination set.
###### Params:
destination, string, Destination set.  
keys, string(variable arguments), Cache keys.
###### Returns:
res, number(int64), Returns the number of elements in the result set on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">set</font>
###### Language: lua
###### Describe:
Set a string value.
###### Params:
key, string, Cache key.  
value, any, Value.  
ttl, number(int), TTL in seconds, <= 0 means persistent.
###### Returns:
res, string or nil, Returns OK on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">setbit</font>
###### Language: lua
###### Describe:
Set the bit value at a specific offset in a string key.
###### Params:
key, string, Cache key.  
offset, number(int64), Offset.  
value, number(int), Bit value to set.
###### Returns:
res, number(int64) or nil, Returns the original bit value at the specified offset on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">setnx</font>
###### Language: lua
###### Describe:
Set a string value only if the key does not exist.
###### Params:
key, string, Cache key.  
value, any, Value.  
ttl, number(int), TTL in seconds, <= 0 means persistent.
###### Returns:
res, bool, Returns true on success, false on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">setxx</font>
###### Language: lua
###### Describe:
Set a string value only if the key exists.
###### Params:
key, string, Cache key.  
value, any, Value.  
ttl, number(int), TTL in seconds, <= 0 means persistent.
###### Returns:
res, bool, Returns true on success, false on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">sinter</font>
###### Language: lua
###### Describe:
Get the intersection of multiple sets.
###### Params:
keys, string(variable arguments), Cache keys.
###### Returns:
res, table(1d), Returns the list of elements in the intersection on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">sinterstore</font>
###### Language: lua
###### Describe:
Compute the intersection of multiple sets and store in the destination set.
###### Params:
destination, string, Destination set.  
keys, string(variable arguments), Cache keys.
###### Returns:
res, number(int64), Returns the number of elements in the result set on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">sismember</font>
###### Language: lua
###### Describe:
Check if a member is in a set.
###### Params:
key, string, Cache key.  
member, string, Member to check.
###### Returns:
res, bool, Returns true if member exists, false otherwise, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">smembers</font>
###### Language: lua
###### Describe:
Get all members of a set.
###### Params:
key, string, Cache key.
###### Returns:
res, table(1d), Returns the list of set members on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">smove</font>
###### Language: lua
###### Describe:
Move a member from one set to another.
###### Params:
source, string, Source key.  
destination, string, Destination key.  
member, string, Member to move.
###### Returns:
res, bool or nil, Returns true if the member was moved, false if not a member of the source set, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">spop</font>
###### Language: lua
###### Describe:
Remove and return a random member from a set.
###### Params:
key, string, Cache key.
###### Returns:
res, string or nil, Returns the removed member on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">spopn</font>
###### Language: lua
###### Describe:
Remove and return a specified number of random members from a set.
###### Params:
key, string, Cache key.  
count, number, Number of members to pop.
###### Returns:
res, table(1d) or nil, Returns the list of removed members on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">srandmember</font>
###### Language: lua
###### Describe:
Get a random member from a set.
###### Params:
key, string, Cache key.
###### Returns:
res, string or nil, Returns a random member on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">srandmembern</font>
###### Language: lua
###### Describe:
Get a specified number of random members from a set.
###### Params:
key, string, Cache key.  
count, number, Number of members.
###### Returns:
res, table(1d) or nil, Returns the list of random members on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">srem</font>
###### Language: lua
###### Describe:
Remove specified members from a set.
###### Params:
key, string, Cache key.  
members, string(variable arguments), Members to remove (multiple can be provided).
###### Returns:
res, number(int64) or nil, Returns the number of removed members on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">sscan</font>
###### Language: lua
###### Describe:
Scan elements in a set.
###### Params:
key, string, Cache key.  
cursor, number(int64), Scan cursor.  
match, string, Match pattern.  
count, number(int64), Expected number of elements to return.
###### Returns:
keys, table(1d) or nil, Returns the scanned elements on success, returns nil on failure.  
cursor, number(int64) or nil, Returns the cursor for the next scan on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">sunion</font>
###### Language: lua
###### Describe:
Get the union of multiple sets.
###### Params:
keys, string(variable arguments), Cache keys.
###### Returns:
res, table(1d), Returns the list of elements in the union on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">sunionstore</font>
###### Language: lua
###### Describe:
Compute the union of multiple sets and store in the destination set.
###### Params:
destination, string, Destination set.  
keys, string(variable arguments), Cache keys.
###### Returns:
res, number(int64), Returns the number of elements in the result set on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">time</font>
###### Language: lua
###### Describe:
Get the server time.
###### Params:
None.
###### Returns:
res, string or nil, Returns the server time (yyyy-MM-dd HH:mm:ss) on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">xadd</font>
###### Language: lua
###### Describe:
Add a message to a stream.
###### Params:
xaddargs, table, Parameters {stream=..., max_len=..., id=..., values=...}.
###### Returns:
res, string or nil, Returns the message ID on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zadd</font>
###### Language: lua
###### Describe:
Add one or more elements to a sorted set.
###### Params:
key, string, Cache key.  
members, table(k/v)(variable arguments), Elements to add {member="mymember", score=123}, multiple can be provided.
###### Returns:
res, number(int64) or nil, Returns the number of new elements added on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zcard</font>
###### Language: lua
###### Describe:
Get the cardinality of a sorted set.
###### Params:
key, string, Cache key.
###### Returns:
res, number(int64), Returns the number of elements in the sorted set on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zcount</font>
###### Language: lua
###### Describe:
Count elements in a sorted set with scores within a range.
###### Params:
key, string, Cache key.  
min, string, Minimum score.  
max, string, Maximum score.
###### Returns:
res, number(int64) or nil, Returns the number of elements in the score range on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zincrby</font>
###### Language: lua
###### Describe:
Increment the score of a member in a sorted set.
###### Params:
key, string, Cache key.  
increment, number, Score increment.  
member, string, Member.
###### Returns:
res, number(float64) or nil, Returns the new score of the member on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zinterstore</font>
###### Language: lua
###### Describe:
Compute the intersection of multiple sorted sets and store in the destination set.
###### Params:
destination, string, Destination set.  
store, table, Cache keys.
###### Returns:
res, number(int64) or nil, Returns the number of elements in the result set on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zlexcount</font>
###### Language: lua
###### Describe:
Count elements in a sorted set within a lexicographic range.
###### Params:
key, string, Cache key.  
min, string, Minimum lexicographic bound, '-' means minimum score, '[min' means inclusive min, '(min' means exclusive min.  
max, string, Maximum lexicographic bound, '+' means maximum score, '[max' means inclusive max, '(max' means exclusive max.
###### Returns:
res, number(int64) or nil, Returns the number of elements in the lexicographic range on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zrange</font>
###### Language: lua
###### Describe:
Get elements from a sorted set by index range (ascending score).
###### Params:
key, string, Cache key.  
start, number, Start index, 0 is the first element.  
end, number, End index, inclusive, -1 is the last element.
###### Returns:
res, table(1d) or nil, Returns the list of elements on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zrangebylex</font>
###### Language: lua
###### Describe:
Get elements from a sorted set by lexicographic range.
###### Params:
key, string, Cache key.  
opt, table, Lexicographic min and max bounds and limit control, e.g. {min="(1", max="[10", offset=10, count=20}.
###### Returns:
res, table(1d) or nil, Returns the list of elements on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zrangebyscore</font>
###### Language: lua
###### Describe:
Get elements from a sorted set by score range.
###### Params:
key, string, Cache key.  
opt, table, Score min and max bounds and limit control, e.g. {min="(1", max="[10", offset=10, count=20}.
###### Returns:
res, table(1d) or nil, Returns the list of elements on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zrangebyscorewithscores</font>
###### Language: lua
###### Describe:
Get elements and their scores from a sorted set by score range.
###### Params:
key, string, Cache key.  
opt, table, Score min and max bounds and limit control, e.g. {min="(1", max="[10", offset=10, count=20}.
###### Returns:
res, table(1d) or nil, Returns the list of elements with scores on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zrangewithscores</font>
###### Language: lua
###### Describe:
Get elements and their scores from a sorted set by index range (ascending score).
###### Params:
key, string, Cache key.  
start, number, Start index, 0 is the first element.  
end, number, End index, inclusive, -1 is the last element.
###### Returns:
res, table(1d) or nil, Returns the list of elements with scores on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zrank</font>
###### Language: lua
###### Describe:
Get the rank of a member in a sorted set (ascending score).
###### Params:
key, string, Cache key.  
member, string, Member.
###### Returns:
res, number(int64) or nil, Returns the rank on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zrem</font>
###### Language: lua
###### Describe:
Remove specified members from a sorted set.
###### Params:
key, string, Cache key.  
members, string(variable arguments), Members to remove (multiple can be provided).
###### Returns:
res, number(int64) or nil, Returns the number of removed members on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zremrangebylex</font>
###### Language: lua
###### Describe:
Remove elements from a sorted set by lexicographic range.
###### Params:
key, string, Cache key.  
min, string, Minimum lexicographic bound.  
max, string, Maximum lexicographic bound.
###### Returns:
res, number(int64) or nil, Returns the number of removed elements on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zremrangebyrank</font>
###### Language: lua
###### Describe:
Remove elements from a sorted set by index range.
###### Params:
key, string, Cache key.  
start, number, Start index.  
stop, number, Stop index.
###### Returns:
res, number(int64) or nil, Returns the number of removed elements on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zremrangebyscore</font>
###### Language: lua
###### Describe:
Remove elements from a sorted set by score range.
###### Params:
key, string, Cache key.  
min, string, Minimum score.  
max, string, Maximum score.
###### Returns:
res, number(int64) or nil, Returns the number of removed elements on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zrevrange</font>
###### Language: lua
###### Describe:
Get elements from a sorted set by index range (descending score).
###### Params:
key, string, Cache key.  
start, number, Start index, 0 is the first element.  
end, number, End index, inclusive, -1 is the last element.
###### Returns:
res, table(1d) or nil, Returns the list of elements on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zrevrangewithscores</font>
###### Language: lua
###### Describe:
Get elements and their scores from a sorted set by index range (descending score).
###### Params:
key, string, Cache key.  
start, number, Start index, 0 is the first element.  
end, number, End index, inclusive, -1 is the last element.
###### Returns:
res, table(1d) or nil, Returns the list of elements with scores on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zrevrank</font>
###### Language: lua
###### Describe:
Get the rank of a member in a sorted set (descending score).
###### Params:
key, string, Cache key.  
member, string, Member.
###### Returns:
res, number(int64) or nil, Returns the rank on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zscan</font>
###### Language: lua
###### Describe:
Scan elements in a sorted set.
###### Params:
key, string, Cache key.  
cursor, number(int64), Scan cursor.  
match, string, Match pattern.  
count, number(int64), Expected number of elements to return.
###### Returns:
keys, table(1d) or nil, Returns the scanned elements on success, returns nil on failure.  
cursor, number(int64) or nil, Returns the cursor for the next scan on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zscore</font>
###### Language: lua
###### Describe:
Get the score of a member in a sorted set.
###### Params:
key, string, Cache key.  
member, string, Member.
###### Returns:
res, number(float64) or nil, Returns the score on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">zunionstore</font>
###### Language: lua
###### Describe:
Compute the union of multiple sorted sets and store in the destination set.
###### Params:
destination, string, Destination set.  
store, table, Cache keys.
###### Returns:
res, number(int64) or nil, Returns the number of elements in the result set on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
#### <font color="DeepSkyBlue">open</font>
##### Language: lua
##### Describe:
Open a Redis connection.  
redis supports connection reuse when calling open on an existing connection.
##### Params:
addr, string, Address.  
password, string, Password.  
db, number(int), Database.
##### Returns:
client, table, Client instance.
##### Since: 0.1
### <font color="DeepSkyBlue">regexps</font>
#### <font color="DeepSkyBlue">find_all_string</font>
##### Language: lua
##### Describe:
Find all matching substrings.
##### Params:
str, string, String to search.  
pattern, string, Regex pattern.
##### Returns:
res, table(1d), Match results.
#### <font color="DeepSkyBlue">find_all_string_submatch</font>
##### Language: lua
##### Describe:
Find all submatches.
##### Params:
str, string, String to search.  
pattern, string, Regex pattern.
##### Returns:
res, table(2d), Match results.
#### <font color="DeepSkyBlue">find_string</font>
##### Language: lua
##### Describe:
Find the first matching substring.
##### Params:
str, string, String to search.  
pattern, string, Regex pattern.
##### Returns:
res, string, Match result.
#### <font color="DeepSkyBlue">find_string_submatch</font>
##### Language: lua
##### Describe:
Find the first submatch.
##### Params:
str, string, String to search.  
pattern, string, Regex pattern.
##### Returns:
res, table(1d), Match result.
#### <font color="DeepSkyBlue">match_string</font>
##### Language: lua
##### Describe:
Check if a string matches a regex pattern.
##### Params:
str, string, String to match.  
pattern, string, Regex pattern.
##### Returns:
res, bool, Match result.
#### <font color="DeepSkyBlue">replace_all_string</font>
##### Language: lua
##### Describe:
Replace all occurrences matching the pattern.
##### Params:
str, string, String to process.  
pattern, string, Regex pattern.  
replace, string, Replacement string.
##### Returns:
res, string, Result string.
#### <font color="DeepSkyBlue">split</font>
##### Language: lua
##### Describe:
Split a string using a regex pattern.
##### Params:
str, string, String to split.  
pattern, string, Regex pattern.
##### Returns:
res, table(1d), List of split strings.
### <font color="DeepSkyBlue">rocketmq</font>
#### <font color="DeepSkyBlue">client</font>
##### <font color="DeepSkyBlue">consumer</font>
###### <font color="DeepSkyBlue">ack_message</font>
###### Language: lua
###### Describe:
Acknowledge a message.
###### Params:
receiptHandle, string, Message receipt handle (msg.ReceiptHandle).  
###### Returns:
err, string or nil, ACK error message, returns nil on success, returns error message string on failure.
###### <font color="DeepSkyBlue">consume_message</font>
###### Language: lua
###### Describe:
Consume messages.
###### Params:
numOfMessages, number(int32), Maximum number of messages to consume at once (up to 16).  
waitseconds, number(int64), Long polling wait time in seconds (up to 30s).  
msgCallback, function(msg), Message callback function (msg is a table).  
errCallback, function(err), Error callback function.  
###### Returns:
None.
##### <font color="DeepSkyBlue">get_consumer</font>
###### Language: lua
###### Describe:
Get a RocketMQ consumer.
###### Params:
instanceId, string, RocketMQ instance ID.  
topic, string, Message topic.  
groupId, string, Consumer group.  
messageTag, string, Message tag.
###### Returns:
consumer, table, Consumer instance.
##### <font color="DeepSkyBlue">get_producer</font>
###### Language: lua
###### Describe:
Get a RocketMQ producer.
###### Params:
instanceId, string, RocketMQ instance ID.  
topic, string, Message topic.
###### Returns:
producer, table, Producer instance.
##### <font color="DeepSkyBlue">producer</font>
###### <font color="DeepSkyBlue">publish_message</font>
###### Language: lua
###### Describe:
Publish a message.
###### Params:
messageBody, string, Message body.  
messageTag, string, Message tag.  
properties, table(k/v), Message properties (both keys and values are strings).  
###### Returns:
messageId, string or nil, Message ID, returns message ID on success, returns nil on failure.  
err, string or nil, Publish error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">open</font>
##### Language: lua
##### Describe:
Create a RocketMQ client.
##### Params:
endpoint, string, Alibaba Cloud endpoint.  
accessKeyId, string, Alibaba Cloud accessKeyId.  
accessKeySecret, string, Alibaba Cloud accessKeySecret.  
securityToken, string, Alibaba Cloud securityToken.
##### Returns:
client, table or nil, Client instance, returns client instance on success, returns nil on failure.
### <font color="DeepSkyBlue">sorts</font>
#### <font color="DeepSkyBlue">numbers</font>
##### Language: lua
##### Describe:
Sort a table of numbers.
##### Params:
arr, table(1d), Table(1d) of numbers to sort.  
asc, bool, Whether to sort in ascending order, optional, default true.
##### Returns:
res, table(1d), Sorted result.
#### <font color="DeepSkyBlue">slice</font>
##### Language: lua
##### Describe:
Sort a table(1d) using a custom comparison function.
##### Params:
arr, table(1d), Table(1d) to sort.  
fun, func(lv, rv) bool, comparison function.
##### Returns:
res, table(1d), Sorted result.
#### <font color="DeepSkyBlue">strings</font>
##### Language: lua
##### Describe:
Sort a table of strings.
##### Params:
arr, table(1d), Table(1d) of strings to sort.  
asc, bool, Whether to sort in ascending order, optional, default true.
##### Returns:
res, table(1d), Sorted result.
### <font color="DeepSkyBlue">sqldbm</font>
#### <font color="DeepSkyBlue">client</font>
##### <font color="DeepSkyBlue">close</font>
###### Language: lua
###### Describe:
Close the SQL data source connection.  
For connection reuse, it is recommended not to close open SQL connections unless necessary.
###### Params:
None.
###### Returns:
err, nil or string, Close connection error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">execute</font>
###### Language: lua
###### Describe:
Execute a SQL query with template parameters.
###### Params:
sqlTpl, string, SQL template to execute.  
sqlArgs, table(k/v), SQL template parameters.  
opts, table(k/v), Optional query parameters, optional, default nil.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">execute_as_rows</font>
###### Language: lua
###### Describe:
Execute a SQL query with template parameters (returns row-oriented table, unlike execute which returns column-oriented).
###### Params:
sqlTpl, string, SQL template to execute.  
sqlArgs, table(k/v), SQL template parameters.  
opts, table(k/v), Optional query parameters, optional, default nil.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">ping</font>
###### Language: lua
###### Describe:
Ping the SQL data source connection.
###### Params:
None.
###### Returns:
err, nil or string, Connection error message, returns nil if connection is normal, returns error message if abnormal.
###### Since: 0.1
##### <font color="DeepSkyBlue">query</font>
###### Language: lua
###### Describe:
Execute a SQL query.
###### Params:
sql, string, SQL statement to execute.  
opts, table(k/v), Optional query parameters, optional, default nil.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">query_as_rows</font>
###### Language: lua
###### Describe:
Execute a SQL query (returns row-oriented table, unlike query which returns column-oriented).
###### Params:
sql, string, SQL statement to execute.  
opts, table(k/v), Optional query parameters, optional, default nil.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
#### <font color="DeepSkyBlue">escape</font>
##### Language: lua
##### Describe:
Escape a string for SQL.  
Currently only supports mysql, other providers return the original string.
##### Params:
provider, string, SQL data source type, supports mysql/mssql/oracle/postgresql/sqlite/clickhouse/elasticsearch/doris/aliyunsls.  
str, string, String to escape.
##### Returns:
res, string, Escaped string.
##### Since: 0.1
#### <font color="DeepSkyBlue">migrate</font>
##### Language: lua
##### Describe:
Data migration.  
Currently only supports mysql to mysql.
##### Params:
srcProvider, string, Source data source type, supports mysql/mssql/oracle/postgresql/sqlite/clickhouse/elasticsearch/doris/aliyunsls.  
srcUrl, string, Source data source connection string.  
srcTable, string, Source table.  
dstProvider, string, Destination data source type, supports mysql/mssql/oracle/postgresql/sqlite/clickhouse/elasticsearch/doris/aliyunsls.  
dstUrl, string, Destination data source connection string.  
dstTable, string, Destination table.  
batchSize, number(int), Batch size for synchronization.
##### Returns:
err, nil or string, Migration error message, returns nil on success, returns error message string on failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">open</font>
##### Language: lua
##### Describe:
Open a SQL data source connection.  
sqldbm supports connection reuse when calling open on an existing connection.
##### Params:
provider, string, SQL data source type, supports mysql/mssql/oracle/postgresql/sqlite/clickhouse/elasticsearch/doris/aliyunsls/duckdb/trino/volcenginetls/filesql.  
url, string, Data source connection string.  
opts, table(k/v), Optional connection parameters, optional, default nil.
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">render</font>
##### Language: lua
##### Describe:
Render a SQL template.
##### Params:
tpl, string, SQL template.  
args, table(k/v), Template parameters.
##### Returns:
res, string, Rendered SQL.
##### Since: 0.1
### <font color="DeepSkyBlue">strings</font>
#### <font color="DeepSkyBlue">contains</font>
##### Language: lua
##### Describe:
Check if a string contains a substring.
##### Params:
str, string, String to check.  
substr, string, Substring.
##### Returns:
res, bool, Result.
#### <font color="DeepSkyBlue">count</font>
##### Language: lua
##### Describe:
Count occurrences of a substring.
##### Params:
str, string, String.  
substr, string, Substring.
##### Returns:
res, number(int), Number of occurrences.
#### <font color="DeepSkyBlue">gb18030_2_utf8</font>
##### Language: lua
##### Describe:
Convert a GB18030 encoded string to UTF-8.
##### Params:
str, string, String to convert.
##### Returns:
res, string or nil, Returns the UTF-8 string on success, returns nil on error.  
err, nil or string, Returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">has_prefix</font>
##### Language: lua
##### Describe:
Check if a string has a prefix.
##### Params:
str, string, String to check.  
prefix, string, Prefix.
##### Returns:
res, bool, Result.
#### <font color="DeepSkyBlue">has_suffix</font>
##### Language: lua
##### Describe:
Check if a string has a suffix.
##### Params:
str, string, String to check.  
suffix, string, Suffix.
##### Returns:
res, bool, Result.
#### <font color="DeepSkyBlue">hash_code</font>
##### Language: lua
##### Describe:
Calculate the hash code of a string.
##### Params:
str, string, String to hash.
##### Returns:
res, string, Hash code string.
#### <font color="DeepSkyBlue">join</font>
##### Language: lua
##### Describe:
Join strings with a separator.
##### Params:
arr, table(1d), List of strings to join.  
sep, string, Separator.
##### Returns:
res, string, Joined string.
#### <font color="DeepSkyBlue">replace</font>
##### Language: lua
##### Describe:
Replace substrings in a string.
##### Params:
str, string, String to process.  
old, string, Old substring.  
new, string, New substring.  
n, number(int), Number of replacements, optional, default -1 replaces all.
##### Returns:
res, string, Resulting string.
#### <font color="DeepSkyBlue">split</font>
##### Language: lua
##### Describe:
Split a string by a separator.
##### Params:
str, string, String to split.  
sep, string, Separator.
##### Returns:
res, table(1d), List of split strings.
#### <font color="DeepSkyBlue">trim</font>
##### Language: lua
##### Describe:
Trim characters from both ends of a string.
##### Params:
str, string, String to trim.  
cutset, string, Set of characters to trim.
##### Returns:
res, string, Trimmed string.
#### <font color="DeepSkyBlue">trim_left</font>
##### Language: lua
##### Describe:
Trim characters from the left side of a string.
##### Params:
str, string, String to trim.  
cutset, string, Set of characters to trim.
##### Returns:
res, string, Trimmed string.
#### <font color="DeepSkyBlue">trim_prefix</font>
##### Language: lua
##### Describe:
Trim a prefix from a string.
##### Params:
str, string, String to trim.  
prefix, string, Prefix to trim.
##### Returns:
res, string, Trimmed string.
#### <font color="DeepSkyBlue">trim_right</font>
##### Language: lua
##### Describe:
Trim characters from the right side of a string.
##### Params:
str, string, String to trim.  
cutset, string, Set of characters to trim.
##### Returns:
res, string, Trimmed string.
#### <font color="DeepSkyBlue">trim_suffix</font>
##### Language: lua
##### Describe:
Trim a suffix from a string.
##### Params:
str, string, String to trim.  
prefix, string, Suffix to trim.
##### Returns:
res, string, Trimmed string.
### <font color="DeepSkyBlue">systems</font>
#### <font color="DeepSkyBlue">environ</font>
##### Language: lua
##### Describe:
Get system environment variables.
##### Params:
None.
##### Returns:
res, table(1d), List of environment variable strings.
### <font color="DeepSkyBlue">templates</font>
#### <font color="DeepSkyBlue">render_to_html</font>
##### Language: lua
##### Describe:
Render data using a template and output to an HTML file.
##### Params:
tpl, string, Template file path.  
data, table, Data to render.  
dst, string, HTML file path.
##### Returns:
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
#### <font color="DeepSkyBlue">render_to_string</font>
##### Language: lua
##### Describe:
Render data using a template and return as a string.
##### Params:
tpl, string, Template file path.  
data, table, Data to render.
##### Returns:
res, string or nil, Call result data, returns rendered string on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
### <font color="DeepSkyBlue">times</font>
#### <font color="DeepSkyBlue">add</font>
##### Language: lua
##### Describe:
Get a time offset by a certain amount from a given time.
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.  
year, number(int), Year offset.  
month, number(int), Month offset.  
day, number(int), Day offset.  
hour, number(int), Hour offset.  
minute, number(int), Minute offset.  
second, number(int), Second offset.
##### Returns:
res, string, Time in yyyy-MM-dd HH:mm:ss format.
#### <font color="DeepSkyBlue">date</font>
##### Language: lua
##### Describe:
Get the date part from a time string.
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, string, Date in yyyy-MM-dd format.
#### <font color="DeepSkyBlue">day</font>
##### Language: lua
##### Describe:
Get the day from a time string.
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, number(int), Day.
#### <font color="DeepSkyBlue">format</font>
##### Language: lua
##### Describe:
Format a time string.
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.  
pattern, string, Format pattern.
##### Returns:
res, string, Formatted time.
#### <font color="DeepSkyBlue">hour</font>
##### Language: lua
##### Describe:
Get the hour from a time string.
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, number(int), Hour.
#### <font color="DeepSkyBlue">minute</font>
##### Language: lua
##### Describe:
Get the minute from a time string.
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, number(int), Minute.
#### <font color="DeepSkyBlue">month</font>
##### Language: lua
##### Describe:
Get the month from a time string.
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, number(int), Month.
#### <font color="DeepSkyBlue">new</font>
##### Language: lua
##### Describe:
Create a time.
##### Params:
year, number(int), Year.  
month, number(int), Month.  
day, number(int), Day.  
hour, number(int), Hour.  
minute, number(int), Minute.  
second, number(int), Second.
##### Returns:
res, string, Time in yyyy-MM-dd HH:mm:ss format.
#### <font color="DeepSkyBlue">now</font>
##### Language: lua
##### Describe:
Get the current time in yyyy-MM-dd HH:mm:ss format.
##### Params:
None.
##### Returns:
res, string, Current time in yyyy-MM-dd HH:mm:ss format.
#### <font color="DeepSkyBlue">now_unix</font>
##### Language: lua
##### Describe:
Get the current Unix timestamp in seconds.
##### Params:
None.
##### Returns:
res, number(int64), Unix timestamp (seconds since 1970-01-01 UTC).
#### <font color="DeepSkyBlue">now_unix_micro</font>
##### Language: lua
##### Describe:
Get the current Unix timestamp in microseconds.
##### Params:
None.
##### Returns:
res, number(int64), Unix timestamp in microseconds.
#### <font color="DeepSkyBlue">now_unix_milli</font>
##### Language: lua
##### Describe:
Get the current Unix timestamp in milliseconds.
##### Params:
None.
##### Returns:
res, number(int64), Unix timestamp in milliseconds.
#### <font color="DeepSkyBlue">now_unix_nano</font>
##### Language: lua
##### Describe:
Get the current Unix timestamp in nanoseconds.
##### Params:
None.
##### Returns:
res, number(int64), Unix timestamp in nanoseconds.
#### <font color="DeepSkyBlue">quarter</font>
##### Language: lua
##### Describe:
Get the quarter from a time string.
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, number(int), Quarter.
#### <font color="DeepSkyBlue">second</font>
##### Language: lua
##### Describe:
Get the second from a time string.
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, number(int), Second.
#### <font color="DeepSkyBlue">sleep</font>
##### Language: lua
##### Describe:
Sleep for a specified number of milliseconds.
##### Params:
mills, number(int64), Milliseconds.
##### Returns:
None.
#### <font color="DeepSkyBlue">sub</font>
##### Language: lua
##### Describe:
Calculate the time difference in seconds between two times.
##### Params:
beg, string, Start time in yyyy-MM-dd HH:mm:ss format.  
end, string, End time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, number(float64), Time difference in seconds.
#### <font color="DeepSkyBlue">to_unix</font>
##### Language: lua
##### Describe:
Convert a time string to Unix timestamp (seconds).
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, number(int), Unix timestamp in seconds.
#### <font color="DeepSkyBlue">to_utc</font>
##### Language: lua
##### Describe:
Convert a time string to UTC format (yyyy-MM-ddTHH:mm:ssZ).
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, string, UTC time in yyyy-MM-ddTHH:mm:ssZ format.
#### <font color="DeepSkyBlue">unix</font>
##### Language: lua
##### Describe:
Create a time from a Unix timestamp.
##### Params:
sec, number(int), Seconds.
##### Returns:
res, string, Time in yyyy-MM-dd HH:mm:ss format.
#### <font color="DeepSkyBlue">utc</font>
##### Language: lua
##### Describe:
Create a time from a UTC string.
##### Params:
tm, string, UTC time in yyyy-MM-ddTHH:mm:ssZ format.
##### Returns:
res, string, Time in yyyy-MM-dd HH:mm:ss format.
#### <font color="DeepSkyBlue">week</font>
##### Language: lua
##### Describe:
Get the week number from a time string.
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, number(int), Week number.
#### <font color="DeepSkyBlue">weekday</font>
##### Language: lua
##### Describe:
Get the day of the week from a time string.
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, number(int), Day of the week.
#### <font color="DeepSkyBlue">year</font>
##### Language: lua
##### Describe:
Get the year from a time string.
##### Params:
tm, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, number(int), Year.
### <font color="DeepSkyBlue">timeseries</font>
#### <font color="DeepSkyBlue">new</font>
##### Language: lua
##### Describe:
Create a time series.
##### Params:
ts, table(1d), Time dimension data (int).  
vs, table(1d), Value dimension data (any).
##### Returns:
series, table, Created time series.
#### <font color="DeepSkyBlue">series</font>
##### <font color="DeepSkyBlue">aggregate</font>
###### Language: lua
###### Describe:
Aggregate a time series.
###### Params:
f, function(ts, vs) any, aggregation function.
###### Returns:
res, any, Aggregation result.
##### <font color="DeepSkyBlue">all</font>
###### Language: lua
###### Describe:
Check if all elements in a time series satisfy a condition.
###### Params:
f, function(i, tick, value) bool, predicate function.
###### Returns:
res, bool, Result.
##### <font color="DeepSkyBlue">any</font>
###### Language: lua
###### Describe:
Check if any element in a time series satisfies a condition.
###### Params:
f, function(i, tick, value) bool, predicate function.
###### Returns:
res, bool, Result.
##### <font color="DeepSkyBlue">beg_tick</font>
###### Language: lua
###### Describe:
Get the first timestamp of a time series.
###### Params:
None.
###### Returns:
begTick, number(int), First timestamp.
##### <font color="DeepSkyBlue">beg_tick_value</font>
###### Language: lua
###### Describe:
Get the first timestamp and value of a time series.
###### Params:
None.
###### Returns:
begTick, number(int), First timestamp.  
begValue, any, First value.
##### <font color="DeepSkyBlue">concat</font>
###### Language: lua
###### Describe:
Concatenate two time series.
###### Params:
other, table, Time series to append.
###### Returns:
series, table, Concatenated time series.
##### <font color="DeepSkyBlue">end_tick</font>
###### Language: lua
###### Describe:
Get the last timestamp of a time series.
###### Params:
None.
###### Returns:
endTick, number(int), Last timestamp.
##### <font color="DeepSkyBlue">end_tick_value</font>
###### Language: lua
###### Describe:
Get the last timestamp and value of a time series.
###### Params:
None.
###### Returns:
endTick, number(int), Last timestamp.  
endValue, any, Last value.
##### <font color="DeepSkyBlue">filter</font>
###### Language: lua
###### Describe:
Filter a time series.
###### Params:
f, function(i, tick, value) bool, filter function.
###### Returns:
series, table, Filtered time series.
##### <font color="DeepSkyBlue">for_each</font>
###### Language: lua
###### Describe:
Iterate over a time series.
###### Params:
f, function(i, tick, value), iteration function.
###### Returns:
None.
##### <font color="DeepSkyBlue">is_empty</font>
###### Language: lua
###### Describe:
Check if a time series is empty.
###### Params:
None.
###### Returns:
empty, bool, True if empty, false otherwise.
##### <font color="DeepSkyBlue">len</font>
###### Language: lua
###### Describe:
Get the length of a time series.
###### Params:
None.
###### Returns:
len, number(int), Length.
##### <font color="DeepSkyBlue">map</font>
###### Language: lua
###### Describe:
Transform a time series using a function.
###### Params:
f, function(i, tick, value) any, transformation function.
###### Returns:
series, table, Transformed time series.
##### <font color="DeepSkyBlue">nearest</font>
###### Language: lua
###### Describe:
Get the nearest timestamp and value in a time series.
###### Params:
tick, number(int), Timestamp.
###### Returns:
tick, number(int), Nearest timestamp.  
value, any, Value at the nearest timestamp.
##### <font color="DeepSkyBlue">quantify</font>
###### Language: lua
###### Describe:
Quantize a time series.
###### Params:
begTick, number(int), Start timestamp for quantization.  
endTick, number(int), End timestamp for quantization.  
step, number(int), Quantization step.
###### Returns:
series, table, Quantized time series.
##### <font color="DeepSkyBlue">sample</font>
###### Language: lua
###### Describe:
Sample a time series.
###### Params:
percent, number(float), Sampling ratio ([0-1]).
###### Returns:
series, table, Sampled time series.
##### <font color="DeepSkyBlue">slice</font>
###### Language: lua
###### Describe:
Slice a time series by timestamp range.
###### Params:
begTick, number(int), Start timestamp.  
endTick, number(int), End timestamp.
###### Returns:
series, table, Sliced time series.
##### <font color="DeepSkyBlue">ticks</font>
###### Language: lua
###### Describe:
Get the timestamps of a time series.
###### Params:
None.
###### Returns:
ts, table(1d), List of timestamps.
##### <font color="DeepSkyBlue">values</font>
###### Language: lua
###### Describe:
Get the values of a time series.
###### Params:
None.
###### Returns:
ts, table(1d), List of values.
### <font color="DeepSkyBlue">types</font>
#### <font color="DeepSkyBlue">duration</font>
##### Language: lua
##### Describe:
Create a duration extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, string, Duration string, e.g., 1ns, 1us, 1s, 1m, 1h, 2h32m13s, etc..
##### Returns:
res, duration, Result value.
#### <font color="DeepSkyBlue">empty_obj</font>
##### Language: lua
##### Describe:
Create an empty_obj extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
None.
##### Returns:
res, empty_obj, Result value.
#### <font color="DeepSkyBlue">float32</font>
##### Language: lua
##### Describe:
Create a float32 extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number, Value.
##### Returns:
res, float32, Result value.
#### <font color="DeepSkyBlue">float64</font>
##### Language: lua
##### Describe:
Create a float64 extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number, Value.
##### Returns:
res, float64, Result value.
#### <font color="DeepSkyBlue">int</font>
##### Language: lua
##### Describe:
Create an int extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number(int), Value.
##### Returns:
res, int, Result value.
#### <font color="DeepSkyBlue">int16</font>
##### Language: lua
##### Describe:
Create an int16 extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number(int), Value.
##### Returns:
res, int16, Result value.
#### <font color="DeepSkyBlue">int32</font>
##### Language: lua
##### Describe:
Create an int32 extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number(int), Value.
##### Returns:
res, int32, Result value.
#### <font color="DeepSkyBlue">int64</font>
##### Language: lua
##### Describe:
Create an int64 extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number(int64), Value.
##### Returns:
res, int64, Result value.
#### <font color="DeepSkyBlue">int8</font>
##### Language: lua
##### Describe:
Create an int8 extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number(int), Value.
##### Returns:
res, int8, Result value.
#### <font color="DeepSkyBlue">null</font>
##### Language: lua
##### Describe:
Create a null extended type, used when a nil value needs to be passed but the Lua table mechanism does not allow it.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
None.
##### Returns:
res, null, Result value.
#### <font color="DeepSkyBlue">time</font>
##### Language: lua
##### Describe:
Create a time extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, string, Time in yyyy-MM-dd HH:mm:ss format.
##### Returns:
res, time, Result value.
#### <font color="DeepSkyBlue">uint</font>
##### Language: lua
##### Describe:
Create a uint extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number(int), Value.
##### Returns:
res, uint, Result value.
#### <font color="DeepSkyBlue">uint16</font>
##### Language: lua
##### Describe:
Create a uint16 extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number(int), Value.
##### Returns:
res, uint16, Result value.
#### <font color="DeepSkyBlue">uint32</font>
##### Language: lua
##### Describe:
Create a uint32 extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number(int), Value.
##### Returns:
res, uint32, Result value.
#### <font color="DeepSkyBlue">uint64</font>
##### Language: lua
##### Describe:
Create a uint64 extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number(int64), Value.
##### Returns:
res, uint64, Result value.
#### <font color="DeepSkyBlue">uint8</font>
##### Language: lua
##### Describe:
Create a uint8 extended type.  
Extended types are only used when some extension functions require fine-grained parameter typing. It is recommended not to use extended types in business logic unless specifically required, to keep code more Lua-compliant.
##### Params:
v, number(int), Value.
##### Returns:
res, uint8, Result value.
### <font color="DeepSkyBlue">urls</font>
#### <font color="DeepSkyBlue">parse</font>
##### Language: lua
##### Describe:
Parse a URL.
##### Params:
url, string, URL string.
##### Returns:
res, table or nil, Returns the parsed URL object on success, returns nil on failure.  
err, nil or string, Returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">path_escape</font>
##### Language: lua
##### Describe:
Escape a URL path.
##### Params:
str, string, String to escape.
##### Returns:
res, string, Escaped string.
#### <font color="DeepSkyBlue">path_unescape</font>
##### Language: lua
##### Describe:
Unescape a URL path.
##### Params:
str, string, String to unescape.
##### Returns:
res, string or nil, Returns the unescaped string on success, returns nil on failure.  
err, nil or string, Returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">query_escape</font>
##### Language: lua
##### Describe:
Escape URL query parameters.
##### Params:
str, string, String to escape.
##### Returns:
res, string, Escaped string.
#### <font color="DeepSkyBlue">query_unescape</font>
##### Language: lua
##### Describe:
Unescape URL query parameters.
##### Params:
str, string, String to unescape.
##### Returns:
res, string or nil, Returns the unescaped string on success, returns nil on failure.  
err, nil or string, Returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">url</font>
##### <font color="DeepSkyBlue">add_value</font>
###### Language: lua
###### Describe:
Add a parameter to the URL.
###### Params:
key, string, Parameter name.  
value, string, Parameter value.
###### Returns:
None.
##### <font color="DeepSkyBlue">get_value</font>
###### Language: lua
###### Describe:
Get a parameter value from the URL.
###### Params:
key, string, Parameter name.
###### Returns:
res, string, Parameter value.
##### <font color="DeepSkyBlue">host</font>
###### Language: lua
###### Describe:
Get the host from the URL.
###### Params:
None.
###### Returns:
res, string, Host information.
##### <font color="DeepSkyBlue">join_path</font>
###### Language: lua
###### Describe:
Append a path to the existing URL path.
###### Params:
path, string, Path string.
###### Returns:
res, table, URL object with the appended path.
##### <font color="DeepSkyBlue">parse</font>
###### Language: lua
###### Describe:
Parse a reference string based on the existing URL.
###### Params:
ref, string, Reference string.
###### Returns:
res, table or nil, Returns the parsed URL object on success, returns nil on failure.  
err, nil or string, Returns nil on success, returns error message on failure.
##### <font color="DeepSkyBlue">path</font>
###### Language: lua
###### Describe:
Get the path from the URL.
###### Params:
None.
###### Returns:
res, string, Path information.
##### <font color="DeepSkyBlue">raw_query</font>
###### Language: lua
###### Describe:
Get the raw query string from the URL.
###### Params:
None.
###### Returns:
res, string, Query string.
##### <font color="DeepSkyBlue">request_uri</font>
###### Language: lua
###### Describe:
Get the request URI from the URL.
###### Params:
None.
###### Returns:
res, string, Request URI.
##### <font color="DeepSkyBlue">schema</font>
###### Language: lua
###### Describe:
Get the schema from the URL.
###### Params:
None.
###### Returns:
res, string, Schema information.
##### <font color="DeepSkyBlue">set_host</font>
###### Language: lua
###### Describe:
Set the host of the URL.
###### Params:
host, string, Host string.
###### Returns:
None.
##### <font color="DeepSkyBlue">set_path</font>
###### Language: lua
###### Describe:
Set the path of the URL.
###### Params:
path, string, Path string.
###### Returns:
None.
##### <font color="DeepSkyBlue">set_schema</font>
###### Language: lua
###### Describe:
Set the schema of the URL.
###### Params:
schema, string, Schema string.
###### Returns:
None.
##### <font color="DeepSkyBlue">set_value</font>
###### Language: lua
###### Describe:
Set a parameter in the URL.
###### Params:
key, string, Parameter name.  
value, string, Parameter value.
###### Returns:
None.
##### <font color="DeepSkyBlue">string</font>
###### Language: lua
###### Describe:
Get the URL string.
###### Params:
None.
###### Returns:
res, string, URL string.
##### <font color="DeepSkyBlue">values</font>
###### Language: lua
###### Describe:
Get all parameters from the URL.
###### Params:
None.
###### Returns:
res, table, All parameters.
### <font color="DeepSkyBlue">volcengine</font>
#### <font color="DeepSkyBlue">open_tos</font>
##### Language: lua
##### Describe:
Open a Volcano Engine TOS (Object Storage Service) connection.
##### Params:
endpoint, string, Volcano Engine endpoint.  
region, string, Volcano Engine region.  
accessKeyId, string, Volcano Engine accessKeyId.  
accessKeySecret, string, Volcano Engine accessKeySecret.
##### Returns:
client, table or nil, Client instance, returns client instance on successful connection, returns nil on connection failure.  
err, nil or string, Connection error message, returns nil on success, returns error message string on connection failure.
##### Since: 0.1
#### <font color="DeepSkyBlue">tos</font>
##### <font color="DeepSkyBlue">copy_object</font>
###### Language: lua
###### Describe:
Copy an object in Volcano Engine TOS.  
Input parameters are as follows:  
params.Bucket, string.  
params.Key, string.  
params.SrcBucket, string.  
params.SrcKey, string.
###### Params:
params, table(k/v), Input parameters.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">delete_object</font>
###### Language: lua
###### Describe:
Delete an object from the specified bucket in Volcano Engine TOS.  
Input parameters are as follows:  
params.Bucket, string.  
params.Key, string.
###### Params:
params, table(k/v), Input parameters.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">does_object_exist</font>
###### Language: lua
###### Describe:
Check whether an object exists in the specified bucket in Volcano Engine TOS.  
Input parameters are as follows:  
params.Bucket, string.  
params.Key, string.
###### Params:
params, table(k/v), Input parameters.
###### Returns:
exist, bool or nil, Returns true or false on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get_object_to_file</font>
###### Language: lua
###### Describe:
Download an object from the specified bucket in Volcano Engine TOS.  
Input parameters are as follows:  
params.Bucket, string.  
params.Key, string.  
params.FilePath, string.
###### Params:
params, table(k/v), Input parameters.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">list_buckets</font>
###### Language: lua
###### Describe:
List buckets in Volcano Engine TOS.  
Input parameters are as follows:  
params.ProjectName, string, project name, optional.  
params.BucketType, string, fns (flat bucket), hns (hierarchical bucket), optional.
###### Params:
params, table(k/v), Input parameters, optional.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">list_objects</font>
###### Language: lua
###### Describe:
List objects in the specified bucket in Volcano Engine TOS.  
Input parameters are as follows:  
params.Bucket, string.  
params.Prefix, string.  
params.Delimiter, string.  
params.StartAfter, string.  
params.ContinuationToken, string.  
params.MaxKeys, number(int).  
params.EncodingType, string.  
params.FetchMeta, bool.  
params.ListOnlyOnce, bool.
###### Params:
params, table(k/v), Input parameters.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">pre_signed_url</font>
###### Language: lua
###### Describe:
Generate a pre-signed URL for an object in Volcano Engine TOS.  
Input parameters are as follows:  
params.HTTPMethod, string.  
params.Bucket, string.  
params.Key, string.  
params.Expires, number(int64).
###### Params:
params, table(k/v), Input parameters.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">put_object_from_file</font>
###### Language: lua
###### Describe:
Upload a file to the specified bucket in Volcano Engine TOS.  
Input parameters are as follows:  
params.Bucket, string.  
params.Key, string.  
params.FilePath, string.
###### Params:
params, table(k/v), Input parameters.
###### Returns:
res, table or nil, Call result data, returns result data on success, returns nil on failure.  
err, nil or string, Call error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
### <font color="DeepSkyBlue">wechat</font>
#### <font color="DeepSkyBlue">client</font>
##### <font color="DeepSkyBlue">access_token</font>
###### Language: lua
###### Describe:
Get the Tencent server access_token.
###### Params:
None.
###### Returns:
res, string or nil, Returns the access_token on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">add_perm_media</font>
###### Language: lua
###### Describe:
Add permanent media.
###### Params:
type, string, Media type: image, voice, video, thumb.  
fn, string, Local path of the media file.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">add_temp_media</font>
###### Language: lua
###### Describe:
Add temporary media.
###### Params:
type, string, Media type: image, voice, video, thumb.  
fn, string, Local path of the media file.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">create_draft</font>
###### Language: lua
###### Describe:
Create a draft.  
Article information includes the following parameters:  
title, string, required, title.  
author, string, optional, author.  
digest, string, optional, summary, only for single-article messages, if not provided, the first 54 characters of the content will be used.  
content, string, required, article content, supports HTML tags, must be less than 20,000 characters and under 1MB, image URLs must come from the 'upload image in article' API, external image URLs will be filtered.  
content_source_url, string, optional, original article URL, i.e., the 'Read More' URL.  
thumb_media_id, string, required, cover image media ID (must be a permanent media ID).  
need_open_comment, number(0 or 1), optional, whether to allow comments, 0 not allowed (default), 1 allowed.  
only_fans_can_comment, number(0 or 1), optional, whether only fans can comment, 0 anyone (default), 1 fans only.
###### Params:
drafts, lua table, Article information, variable arguments (at least one).
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">del_draft</font>
###### Language: lua
###### Describe:
Delete a draft.
###### Params:
media_id, string, Media ID of the draft.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">del_perm_media</font>
###### Language: lua
###### Describe:
Delete permanent media.
###### Params:
media_id, string, Media ID to delete.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">del_publish</font>
###### Language: lua
###### Describe:
Delete a published article.
###### Params:
article_id, string, Article ID returned on successful publish.  
index, number(int), Position of the article in the message, starting from 1, 0 deletes all articles, optional, default 0.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get_access_token</font>
###### Language: lua
###### Describe:
Get the Tencent server access_token.
###### Params:
None.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get_api_domain_ip</font>
###### Language: lua
###### Describe:
Get WeChat API domain IP addresses.
###### Params:
None.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get_article</font>
###### Language: lua
###### Describe:
Get a published article.
###### Params:
article_id, string, Article ID returned on successful publish.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get_callback_ip</font>
###### Language: lua
###### Describe:
Get WeChat callback IP addresses.
###### Params:
None.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get_draft</font>
###### Language: lua
###### Describe:
Get a draft.
###### Params:
media_id, string, Media ID of the draft.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get_draft_list</font>
###### Language: lua
###### Describe:
Get the list of drafts.
###### Params:
offset, number(int), Offset (0-based).  
count, number(int, [1,20]), number to fetch.  
no_content, number(0 or 1), Whether to exclude article content, 0 returns content, 1 excludes content, optional, default 0.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get_perm_media_list</font>
###### Language: lua
###### Describe:
Get the list of permanent media.
###### Params:
type, string, Media type: image, voice, video, thumb.  
offset, number(int), Offset from the start, 0 means the first.  
count, number(int), Number of media to fetch, between 1 and 20.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get_publish_list</font>
###### Language: lua
###### Describe:
Get the list of published articles.
###### Params:
offset, number(int), Offset (0-based).  
count, number(int, [1,20]), number to fetch.  
no_content, number(0 or 1), Whether to exclude article content, 0 returns content, 1 excludes content, optional, default 0.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">get_publish_state</font>
###### Language: lua
###### Describe:
Get the publish state.  
Publish state: 0: success, 1: publishing, 2: originality check failed, 3: general failure, 4: platform review failed, 5: user deleted all articles after success, 6: system banned all articles after success.
###### Params:
publish_id, number(int64), Publish task ID returned when publishing a draft.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
##### <font color="DeepSkyBlue">publish_draft</font>
###### Language: lua
###### Describe:
Publish a draft.  
The publish API is mainly used to generate a permanent article link, unlike manual publishing on the official platform (manual publishing shows on the official account homepage and pushes to followers), API publishing does not show on the homepage nor push to followers.
###### Params:
media_id, string, Media ID of the draft.
###### Returns:
res, lua table or nil, Returns result data on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message string on failure.
###### Since: 0.1
#### <font color="DeepSkyBlue">open</font>
##### Language: lua
##### Describe:
Open a connection to the Tencent server.
##### Params:
appId, string, appId.  
appSecret, string, appSecret.  
token, string, token.  
encodingAesKey, string, encodingAesKey.  
proxy, string, proxy.
##### Returns:
res, table, Client instance.
##### Since: 0.1
### <font color="DeepSkyBlue">xmls</font>
#### <font color="DeepSkyBlue">load</font>
##### Language: lua
##### Describe:
Load and parse an XML file.
##### Params:
path, string, XML file path.
##### Returns:
res, table or nil, Returns a Lua table on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">parse</font>
##### Language: lua
##### Describe:
Parse XML to a table.
##### Params:
content, string, XML string.
##### Returns:
res, table or nil, Returns a Lua table on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">to_json</font>
##### Language: lua
##### Describe:
Convert XML to JSON.
##### Params:
content, string, XML string.
##### Returns:
res, string or nil, Returns the JSON string on success, returns nil on failure.  
err, nil or string, Error message, returns nil on success, returns error message on failure.
### <font color="DeepSkyBlue">yamls</font>
#### <font color="DeepSkyBlue">dump</font>
##### Language: lua
##### Describe:
Convert a table to YAML.
##### Params:
obj, table, Object to serialize.
##### Returns:
res, string, YAML string.
#### <font color="DeepSkyBlue">from_json</font>
##### Language: lua
##### Describe:
Convert JSON to YAML.
##### Params:
content, string, JSON string.
##### Returns:
res, string or nil, Returns the YAML string on success, returns nil on failure.  
err, string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">load</font>
##### Language: lua
##### Describe:
Load and parse a YAML file.
##### Params:
path, string, YAML file path.
##### Returns:
res, table or nil, Returns a Lua table on success, returns nil on failure.  
err, string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">parse</font>
##### Language: lua
##### Describe:
Parse YAML to a table.
##### Params:
content, string, YAML string.
##### Returns:
res, table or nil, Returns a Lua table on success, returns nil on failure.  
err, string, Error message, returns nil on success, returns error message on failure.
#### <font color="DeepSkyBlue">to_json</font>
##### Language: lua
##### Describe:
Convert YAML to JSON.
##### Params:
content, string, YAML string.
##### Returns:
res, string or nil, Returns the JSON string on success, returns nil on failure.  
err, string, Error message, returns nil on success, returns error message on failure.  
