import React, { useState } from 'react';
import { Row, Col, Statistic, Icon, Card } from 'antd';

const Repository = () => {
  return (
    <div style={{ background: '#ECECEC', padding: '30px' }}>
      <Row gutter={8}>
        <Col span={8}>
          <Card>
            <Statistic
              title="Database"
              value={11.28}
              precision={2}
              valueStyle={{ color: '#3f8600' }}
              prefix={<Icon type="arrow-up" />}
              suffix="%"
            />
          </Card>
        </Col>
      </Row>
      <Row gutter={16}>
        <Col span={12}>
          <Card>
            <Statistic
              title="Active"
              value={11.28}
              precision={2}
              valueStyle={{ color: '#3f8600' }}
              prefix={<Icon type="arrow-up" />}
              suffix="%"
            />
          </Card>
        </Col>
        <Col span={12}>
          <Card>
            <Statistic
              title="Idle"
              value={9.3}
              precision={2}
              valueStyle={{ color: '#cf1322' }}
              prefix={<Icon type="arrow-down" />}
              suffix="%"
            />
          </Card>
        </Col>
      </Row>
    </div>
  );
};

export default Repository;
