import React, { useState } from 'react';
import Layout from 'antd/lib/layout';
import Menu from 'antd/lib/menu';
import Icon from 'antd/lib/icon';
import Breadcrumb from 'antd/lib/breadcrumb';
import SubMenu from 'antd/lib/menu/SubMenu';
import Repository from '../respository/Repository';
const { Header, Content, Footer, Sider } = Layout;

const MenuLayout = () => {
  const [collapsed, setCollapse] = useState(false);
  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider
        collapsible
        collapsed={collapsed}
        onCollapse={() => setCollapse(!collapsed)}
      >
        <div className="logo" />
        <Menu theme="dark" defaultSelectedKeys={['1']} mode="inline">
          <Menu.Item key="1">
            <Icon type="desktop" />
            <span>FEGA</span>
          </Menu.Item>
          <SubMenu
            key="sub1"
            title={
              <span>
                <Icon type="user" />
                <span>Middleware</span>
              </span>
            }
          >
            <Menu.Item key="2">Tom</Menu.Item>
            <Menu.Item key="3">Bill</Menu.Item>
            <Menu.Item key="4">Alex</Menu.Item>
          </SubMenu>
          <SubMenu
            key="sub2"
            title={
              <span>
                <Icon type="database" theme="filled" />
                <span>Repository</span>
              </span>
            }
          >
            <Menu.Item key="5">
              <span>
                <Icon type="hdd" theme="filled" />
                <span>Database</span>
              </span>
            </Menu.Item>
            <Menu.Item key="6">
              <span>
                <Icon type="api" theme="filled" />
                <span>External APIs</span>
              </span>
            </Menu.Item>
          </SubMenu>
        </Menu>
      </Sider>
      <Layout>
        <Header style={{ background: '#fff', padding: 0 }} />
        <Content style={{ margin: '0 16px' }}>
          <Breadcrumb style={{ margin: '16px 0' }}>
            <Breadcrumb.Item>Repository</Breadcrumb.Item>
            <Breadcrumb.Item>Database</Breadcrumb.Item>
          </Breadcrumb>
          <div
            style={{
              padding: 24,
              background: '#fff',
              minHeight: 360,
              margin: '16px 0'
            }}
          >
            <Repository />
          </div>
        </Content>
        <Footer style={{ textAlign: 'center' }}>
          Ant Design ©2018 Created by Ant UED
        </Footer>
      </Layout>
    </Layout>
  );
};

export default MenuLayout;