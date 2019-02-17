import React from 'react';
import { default as MenuLayout } from './components/common/MenuLayout';
import Repository from './pages/Repository';
import './App.css';

const App = () => {
  return (
    <MenuLayout>
      <Repository />
    </MenuLayout>
  );
};

export default App;
