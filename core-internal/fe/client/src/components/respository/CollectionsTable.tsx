import React from 'react';
import { Table, Divider, Tag } from 'antd';
import { Query } from 'react-apollo';
import GET_COLLECTIONS from '../../gql/queries/getCollections';

const { Column, ColumnGroup } = Table;

type data = {
  id: string;
  name: string;
  host: string;
};

// const data: data[] = [
//   {
//     key: '1',
//     firstName: 'John',
//     lastName: 'Brown',
//     age: 32,
//     address: 'New York No. 1 Lake Park',
//     tags: ['nice', 'developer']
//   },
//   {
//     key: '2',
//     firstName: 'Jim',
//     lastName: 'Green',
//     age: 42,
//     address: 'London No. 1 Lake Park',
//     tags: ['loser']
//   },
//   {
//     key: '3',
//     firstName: 'Joe',
//     lastName: 'Black',
//     age: 32,
//     address: 'Sidney No. 1 Lake Park',
//     tags: ['cool', 'teacher']
//   }
// ];

const CollectionsTable = () => (
  <Query query={GET_COLLECTIONS}>
    {({ data, loading, error }) => {
      if (loading) return <></>;
      if (error) return <p>ERROR</p>;
      console.log(data);
      return (
        <Table dataSource={data} loading={loading ? true : false}>
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
                <a href="javascript:;">Invite {record.name}</a>
                <Divider type="vertical" />
                <a href="javascript:;">Delete</a>
              </span>
            )}
          />
        </Table>
      );
    }}
  </Query>
);

export default CollectionsTable;
