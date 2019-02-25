import React from 'react';
import { default as MenuLayout } from './components/common/MenuLayout';
import Repository from './pages/Repository';
import ApolloClient from 'apollo-boost';
import { ApolloProvider } from 'react-apollo';
import './App.css';

const client = new ApolloClient({
  uri: 'http://localhost:8080/query'
});

const App = () => {
  return (
    <ApolloProvider client={client}>
      <MenuLayout>
        <Repository />
      </MenuLayout>
    </ApolloProvider>
  );
};

export default App;
