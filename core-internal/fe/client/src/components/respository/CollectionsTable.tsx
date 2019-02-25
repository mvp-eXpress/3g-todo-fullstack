import React from 'react';
import { Table, Divider, Tag } from 'antd';

const { Column, ColumnGroup } = Table;

type data = {
  key: string;
  firstName: string;
  lastName: string;
  age: number;
  address: string;
  tags: string[];
};

const data: data[] = [
  {
    key: '1',
    firstName: 'John',
    lastName: 'Brown',
    age: 32,
    address: 'New York No. 1 Lake Park',
    tags: ['nice', 'developer']
  },
  {
    key: '2',
    firstName: 'Jim',
    lastName: 'Green',
    age: 42,
    address: 'London No. 1 Lake Park',
    tags: ['loser']
  },
  {
    key: '3',
    firstName: 'Joe',
    lastName: 'Black',
    age: 32,
    address: 'Sidney No. 1 Lake Park',
    tags: ['cool', 'teacher']
  }
];

const CollectionsTable = () => (
  <Table dataSource={data}>
    <Column title="First Name" dataIndex="firstName" key="firstName" />
    <Column title="Last Name" dataIndex="lastName" key="lastName" />
    <Column title="Age" dataIndex="age" key="age" />
    <Column title="Address" dataIndex="address" key="address" />
    <Column
      title="Tags"
      dataIndex="tags"
      key="tags"
      render={tags => (
        <span>
          {tags.map((tag: any) => (
            <Tag color="blue" key={tag}>
              {tag}
            </Tag>
          ))}
        </span>
      )}
    />
    <Column
      title="Action"
      key="action"
      render={(text, record: data) => (
        <span>
          <a href="javascript:;">Invite {record.lastName}</a>
          <Divider type="vertical" />
          <a href="javascript:;">Delete</a>
        </span>
      )}
    />
  </Table>
);

export default CollectionsTable;
