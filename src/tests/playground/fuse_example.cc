/**
 * Copyright © 2018 The IPFN Developers. All Rights Reserved.
 * 
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * 
 *     http://www.apache.org/licenses/LICENSE-2.0
 * 
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
#define FUSE_USE_VERSION 26

#include <fuse.h>
#include <stdio.h>
#include <string.h>
#include <errno.h>
#include <fcntl.h>
#include <iosfwd>

#include <glog/logging.h>

static const char *hello_str = "Hello World!\n";
static const char *hello_path = "/hello";

static int hello_getattr(const char *path, struct stat *stbuf)
{
    LOG(INFO) << "getattr: " << path;
    int res = 0;

    memset(stbuf, 0, sizeof(struct stat));
    if (strcmp(path, "/") == 0)
    {
        stbuf->st_mode = S_IFDIR | 0755;
        stbuf->st_nlink = 2;
    }
    else if (strcmp(path, hello_path) == 0)
    {
        stbuf->st_mode = S_IFREG | 0444;
        stbuf->st_nlink = 1;
        stbuf->st_size = strlen(hello_str);
    }
    else
        res = -ENOENT;

    return res;
}

static int hello_readdir(const char *path, void *buf, fuse_fill_dir_t filler,
                         off_t offset, struct fuse_file_info *fi)
{
    LOG(INFO) << "readdir: " << path;
    (void)offset;
    (void)fi;

    if (strcmp(path, "/") != 0)
        return -ENOENT;

    filler(buf, ".", NULL, 0);
    filler(buf, "..", NULL, 0);
    filler(buf, hello_path + 1, NULL, 0);

    return 0;
}

static int hello_open(const char *path, struct fuse_file_info *fi)
{
    LOG(INFO) << "open: " << path;
    if (strcmp(path, hello_path) != 0)
        return -ENOENT;

    // if ((fi->flags & 3) != O_RDONLY)
    //     return -EACCES;

    return 0;
}

static int hello_read(const char *path, char *buf, size_t size, off_t offset,
                      struct fuse_file_info *fi)
{
    LOG(INFO) << "read: " << path;
    size_t len;
    (void)fi;
    if (strcmp(path, hello_path) != 0)
        return -ENOENT;

    len = strlen(hello_str);
    if (offset < (unsigned)len)
    {
        if (offset + size > len)
            size = len - offset;
        memcpy(buf, hello_str + offset, size);
    }
    else
        size = 0;

    return size;
}

static int hello_write(const char *path, const char *buf, size_t size, off_t offset,
                       struct fuse_file_info *fi)
{
    LOG(INFO) << "write: " << path << " content: " << buf;
    return size;
}

static int hello_fsync(const char *path, int sync, struct fuse_file_info *fi)
{
    LOG(INFO) << "fsync: " << path << " sync: " << sync << " fh: " << fi->fh;
    return 0;
}

static void *hello_init(struct fuse_conn_info *conn)
{
    (void)conn;
    return NULL;
}

static int hello_write_buf(const char *path, struct fuse_bufvec *buf, off_t off,
                           struct fuse_file_info *fi)
{
    LOG(INFO) << "write_buf: " << path << " buf: " << buf->buf << " count: " << buf->count;
    return 0;
}

struct hello_fuse_operations : fuse_operations
{
    hello_fuse_operations()
    {
        init = hello_init;
        getattr = hello_getattr;
        readdir = hello_readdir;
        open = hello_open;
        read = hello_read;
        write = hello_write;
        fsync = hello_fsync;

        write_buf = hello_write_buf;
    }
};

static struct hello_fuse_operations hello_oper;

int main(int argc, char *argv[])
{
    google::InitGoogleLogging(argv[0]);
    return fuse_main(argc, argv, &hello_oper, NULL);
}