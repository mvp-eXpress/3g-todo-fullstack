import React from 'react';
import { Avatar, List } from 'antd';
import { Collection } from '../../models/Collection';

const CollectionsList = () => {
  return (
    <>
      <List
        itemLayout="horizontal"
        renderItem={(item: Collection) => (
          <List.Item>
            <List.Item.Meta
              avatar={
                <Avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
              }
              title={<a href="https://ant.design">{item.name}</a>}
              description="Ant Design, a design language for background applications, is refined by Ant UED Team"
            />
          </List.Item>
        )}
      />
    </>
  );
};

export default CollectionsList;
