#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
API 测试脚本
测试所有后端 API 接口
"""

import requests
import json
import time
from typing import Optional, Dict, Any

# 配置
BASE_URL = "http://localhost:8080"
API_BASE = f"{BASE_URL}/api/v1"

# 全局变量，用于存储测试数据
test_data = {
    "token": None,
    "user_id": None,
    "role_id": None,
    "menu_id": None,
    "log_id": None,
    "test_username": f"username{int(time.time())}",
    "test_password": f"password{int(time.time())}",
}


def print_result(test_name: str, status: str, response: Optional[requests.Response] = None, error: Optional[str] = None):
    """打印测试结果"""
    status_icon = "✓" if status == "PASS" else "✗"
    print(f"\n{status_icon} {test_name} - {status}")

    if response:
        try:
            print(f"  状态码: {response.status_code}")
            print(f"  响应: {json.dumps(response.json(), ensure_ascii=False, indent=2)}")
        except:
            print(f"  响应: {response.text}")

    if error:
        print(f"  错误: {error}")


def make_request(method: str, endpoint: str, data: Optional[Dict] = None, headers: Optional[Dict] = None, params: Optional[Dict] = None) -> requests.Response:
    """发送 HTTP 请求"""
    url = f"{API_BASE}{endpoint}"

    if method == "GET":
        return requests.get(url, headers=headers, params=params)
    elif method == "POST":
        return requests.post(url, json=data, headers=headers)
    elif method == "PUT":
        return requests.put(url, json=data, headers=headers)
    elif method == "DELETE":
        return requests.delete(url, headers=headers)
    else:
        raise ValueError(f"不支持的 HTTP 方法: {method}")


def test_captcha():
    """测试验证码接口"""
    try:
        response = make_request("GET", "/captcha")
        if response.status_code == 200:
            print_result("获取验证码", "PASS", response)
            return True
        else:
            print_result("获取验证码", "FAIL", response)
            return False
    except Exception as e:
        print_result("获取验证码", "FAIL", error=str(e))
        return False


def test_register():
    """测试用户注册"""
    try:
        data = {
            "username": test_data["test_username"],
            "password": test_data["test_password"],
            "email": f"{test_data['test_username']}@example.com",
            "nickname": "测试用户"
        }
        response = make_request("POST", "/register", data=data)
        if response.status_code == 200:
            result = response.json()
            test_data["user_id"] = result["data"]["id"]
            print_result("用户注册", "PASS", response)
            return True
        else:
            print_result("用户注册", "FAIL", response)
            return False
    except Exception as e:
        print_result("用户注册", "FAIL", error=str(e))
        return False


def test_login():
    """测试用户登录"""
    try:
        data = {
            "username": test_data["test_username"],
            "password": test_data["test_password"]
        }
        response = make_request("POST", "/login", data=data)
        if response.status_code == 200:
            result = response.json()
            test_data["token"] = result["data"]["token"]
            test_data["user_id"] = result["data"]["user"]["id"]
            print_result("用户登录", "PASS", response)
            return True
        else:
            print_result("用户登录", "FAIL", response)
            return False
    except Exception as e:
        print_result("用户登录", "FAIL", error=str(e))
        return False


def get_auth_headers() -> Dict[str, str]:
    """获取认证头"""
    return {"Authorization": f"Bearer {test_data['token']}"}


def test_get_users():
    """测试获取用户列表"""
    try:
        response = make_request("GET", "/users", headers=get_auth_headers())
        if response.status_code == 200:
            print_result("获取用户列表", "PASS", response)
            return True
        else:
            print_result("获取用户列表", "FAIL", response)
            return False
    except Exception as e:
        print_result("获取用户列表", "FAIL", error=str(e))
        return False


def test_get_user_info():
    """测试获取用户信息"""
    try:
        response = make_request("GET", f"/users/{test_data['user_id']}", headers=get_auth_headers())
        if response.status_code == 200:
            print_result("获取用户信息", "PASS", response)
            return True
        else:
            print_result("获取用户信息", "FAIL", response)
            return False
    except Exception as e:
        print_result("获取用户信息", "FAIL", error=str(e))
        return False


def test_update_user():
    """测试更新用户信息"""
    try:
        data = {
            "username": test_data["test_username"],
            "email": f"{test_data['test_username']}_updated@example.com",
            "nickname": "测试用户（已更新）"
        }
        response = make_request("PUT", f"/users/{test_data['user_id']}", data=data, headers=get_auth_headers())
        if response.status_code == 200:
            print_result("更新用户信息", "PASS", response)
            return True
        else:
            print_result("更新用户信息", "FAIL", response)
            return False
    except Exception as e:
        print_result("更新用户信息", "FAIL", error=str(e))
        return False


def test_update_user_status():
    """测试更新用户状态"""
    try:
        data = {"status": 1}
        response = make_request("PUT", f"/users/{test_data['user_id']}/status", data=data, headers=get_auth_headers())
        if response.status_code == 200:
            print_result("更新用户状态", "PASS", response)
            return True
        else:
            print_result("更新用户状态", "FAIL", response)
            return False
    except Exception as e:
        print_result("更新用户状态", "FAIL", error=str(e))
        return False


def test_change_password():
    """测试修改密码"""
    try:
        data = {
            "old_password": test_data["test_password"],
            "new_password": "NewTest123456"
        }
        response = make_request("PUT", "/users/change-password", data=data, headers=get_auth_headers())
        if response.status_code == 200:
            test_data["test_password"] = data["new_password"]
            print_result("修改密码", "PASS", response)
            return True
        else:
            print_result("修改密码", "FAIL", response)
            return False
    except Exception as e:
        print_result("修改密码", "FAIL", error=str(e))
        return False


def test_create_role():
    """测试创建角色"""
    try:
        data = {
            "name": f"测试角色_{int(time.time())}",
            "description": "这是一个测试角色",
            "status": 1
        }
        response = make_request("POST", "/roles", data=data, headers=get_auth_headers())
        if response.status_code == 200:
            result = response.json()
            test_data["role_id"] = result["data"]["id"]
            print_result("创建角色", "PASS", response)
            return True
        else:
            print_result("创建角色", "FAIL", response)
            return False
    except Exception as e:
        print_result("创建角色", "FAIL", error=str(e))
        return False


def test_get_roles():
    """测试获取角色列表"""
    try:
        response = make_request("GET", "/roles", headers=get_auth_headers())
        if response.status_code == 200:
            print_result("获取角色列表", "PASS", response)
            return True
        else:
            print_result("获取角色列表", "FAIL", response)
            return False
    except Exception as e:
        print_result("获取角色列表", "FAIL", error=str(e))
        return False


def test_get_role():
    """测试获取角色详情"""
    try:
        response = make_request("GET", f"/roles/{test_data['role_id']}", headers=get_auth_headers())
        if response.status_code == 200:
            print_result("获取角色详情", "PASS", response)
            return True
        else:
            print_result("获取角色详情", "FAIL", response)
            return False
    except Exception as e:
        print_result("获取角色详情", "FAIL", error=str(e))
        return False


def test_update_role():
    """测试更新角色"""
    try:
        data = {
            "name": f"测试角色_{int(time.time())}_updated",
            "description": "这是一个更新后的测试角色",
            "status": 1
        }
        response = make_request("PUT", f"/roles/{test_data['role_id']}", data=data, headers=get_auth_headers())
        if response.status_code == 200:
            print_result("更新角色", "PASS", response)
            return True
        else:
            print_result("更新角色", "FAIL", response)
            return False
    except Exception as e:
        print_result("更新角色", "FAIL", error=str(e))
        return False


def test_create_menu():
    """测试创建菜单"""
    try:
        data = {
            "name": f"test_menu_{int(time.time())}",
            "title": "测试菜单",
            "path": "/test",
            "component": "Test",
            "icon": "test-icon",
            "sort": 100,
            "status": 1
        }
        response = make_request("POST", "/menus", data=data, headers=get_auth_headers())
        if response.status_code == 200:
            result = response.json()
            test_data["menu_id"] = result["data"]["id"]
            print_result("创建菜单", "PASS", response)
            return True
        else:
            print_result("创建菜单", "FAIL", response)
            return False
    except Exception as e:
        print_result("创建菜单", "FAIL", error=str(e))
        return False


def test_get_menu_tree():
    """测试获取菜单树"""
    try:
        response = make_request("GET", "/menus", headers=get_auth_headers())
        if response.status_code == 200:
            print_result("获取菜单树", "PASS", response)
            return True
        else:
            print_result("获取菜单树", "FAIL", response)
            return False
    except Exception as e:
        print_result("获取菜单树", "FAIL", error=str(e))
        return False


def test_get_all_menus():
    """测试获取所有菜单"""
    try:
        response = make_request("GET", "/menus/all", headers=get_auth_headers())
        if response.status_code == 200:
            print_result("获取所有菜单", "PASS", response)
            return True
        else:
            print_result("获取所有菜单", "FAIL", response)
            return False
    except Exception as e:
        print_result("获取所有菜单", "FAIL", error=str(e))
        return False


def test_get_menu():
    """测试获取菜单详情"""
    try:
        response = make_request("GET", f"/menus/{test_data['menu_id']}", headers=get_auth_headers())
        if response.status_code == 200:
            print_result("获取菜单详情", "PASS", response)
            return True
        else:
            print_result("获取菜单详情", "FAIL", response)
            return False
    except Exception as e:
        print_result("获取菜单详情", "FAIL", error=str(e))
        return False


def test_update_menu():
    """测试更新菜单"""
    try:
        data = {
            "name": f"test_menu_{int(time.time())}_updated",
            "title": "测试菜单（已更新）",
            "path": "/test-updated",
            "component": "Test",
            "icon": "test-icon-updated",
            "sort": 101,
            "status": 1
        }
        response = make_request("PUT", f"/menus/{test_data['menu_id']}", data=data, headers=get_auth_headers())
        if response.status_code == 200:
            print_result("更新菜单", "PASS", response)
            return True
        else:
            print_result("更新菜单", "FAIL", response)
            return False
    except Exception as e:
        print_result("更新菜单", "FAIL", error=str(e))
        return False


def test_get_operation_logs():
    """测试获取操作日志"""
    try:
        response = make_request("GET", "/operation-logs", headers=get_auth_headers())
        if response.status_code == 200:
            result = response.json()
            if result["data"]["list"]:
                test_data["log_id"] = result["data"]["list"][0]["id"]
            print_result("获取操作日志", "PASS", response)
            return True
        else:
            print_result("获取操作日志", "FAIL", response)
            return False
    except Exception as e:
        print_result("获取操作日志", "FAIL", error=str(e))
        return False


def test_delete_menu():
    """测试删除菜单"""
    try:
        response = make_request("DELETE", f"/menus/{test_data['menu_id']}", headers=get_auth_headers())
        if response.status_code == 200:
            print_result("删除菜单", "PASS", response)
            return True
        else:
            print_result("删除菜单", "FAIL", response)
            return False
    except Exception as e:
        print_result("删除菜单", "FAIL", error=str(e))
        return False


def test_delete_role():
    """测试删除角色"""
    try:
        response = make_request("DELETE", f"/roles/{test_data['role_id']}", headers=get_auth_headers())
        if response.status_code == 200:
            print_result("删除角色", "PASS", response)
            return True
        else:
            print_result("删除角色", "FAIL", response)
            return False
    except Exception as e:
        print_result("删除角色", "FAIL", error=str(e))
        return False


def test_delete_user():
    """测试删除用户"""
    try:
        response = make_request("DELETE", f"/users/{test_data['user_id']}", headers=get_auth_headers())
        if response.status_code == 200:
            print_result("删除用户", "PASS", response)
            return True
        else:
            print_result("删除用户", "FAIL", response)
            return False
    except Exception as e:
        print_result("删除用户", "FAIL", error=str(e))
        return False


def main():
    """主测试函数"""
    print("=" * 60)
    print("API 测试脚本")
    print(f"测试地址: {BASE_URL}")
    print("=" * 60)

    # 检查服务器是否运行
    try:
        response = requests.get(BASE_URL, timeout=5)
        print(f"\n✓ 服务器运行中")
    except:
        print(f"\n✗ 无法连接到服务器 {BASE_URL}")
        print("请确保后端服务已启动")
        return

    # 测试用例列表
    tests = [
        # 公开接口
        ("获取验证码", test_captcha),
        ("用户注册", test_register),
        ("用户登录", test_login),

        # 用户管理
        ("获取用户列表", test_get_users),
        ("获取用户信息", test_get_user_info),
        ("更新用户信息", test_update_user),
        ("更新用户状态", test_update_user_status),
        ("修改密码", test_change_password),

        # 角色管理
        ("创建角色", test_create_role),
        ("获取角色列表", test_get_roles),
        ("获取角色详情", test_get_role),
        ("更新角色", test_update_role),

        # 菜单管理
        ("创建菜单", test_create_menu),
        ("获取菜单树", test_get_menu_tree),
        ("获取所有菜单", test_get_all_menus),
        ("获取菜单详情", test_get_menu),
        ("更新菜单", test_update_menu),

        # 操作日志
        ("获取操作日志", test_get_operation_logs),

        # 清理测试数据
#         ("删除菜单", test_delete_menu),
#         ("删除角色", test_delete_role),
#         ("删除用户", test_delete_user),
    ]

    # 执行测试
    passed = 0
    failed = 0

    for test_name, test_func in tests:
        if test_func():
            passed += 1
        else:
            failed += 1
        time.sleep(0.5)  # 避免请求过快

    # 打印测试结果汇总
    print("\n" + "=" * 60)
    print("测试结果汇总")
    print("=" * 60)
    print(f"总计: {len(tests)} 个测试")
    print(f"通过: {passed} 个")
    print(f"失败: {failed} 个")
    print(f"成功率: {passed/len(tests)*100:.1f}%")
    print("=" * 60)


if __name__ == "__main__":
    main()