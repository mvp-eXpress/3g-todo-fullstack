import React, { ReactPropTypes } from 'react';
import { Avatar, List } from 'antd';
import { Collection } from '../../models/Collection';
import { Query } from 'react-apollo';
import GET_COLLECTIONS from '../../gql/queries/getCollections';

const CollectionsListWithData = () => (
  <Query query={GET_COLLECTIONS}>
    {({ data, loading, error }) => {
      if (loading) return <></>;
      if (error) return <p>ERROR</p>;
      return <CollectionsList data={data} />;
    }}
  </Query>
);

const CollectionsList = (props: { data: any }) => {
  return (
    <List
      itemLayout="horizontal"
      dataSource={props.data.getCollections}
      renderItem={(item: Collection) => <CollectionsListItem item={item} />}
    />
  );
};

const CollectionsListItem = (props: { item: Collection }) => {
  return (
    <List.Item>
      <List.Item.Meta
        avatar={
          <Avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
        }
        title={<a href="https://ant.design">{props.item.name}</a>}
        description={props.item.host}
      />
    </List.Item>
  );
};

export default CollectionsListWithData;
