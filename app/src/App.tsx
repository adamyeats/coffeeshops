import React from 'react';

import { Sidebar } from './components/Sidebar';
import { MapSheet } from './components/MapSheet';

const App: React.FC = () => (
  <div className='App'>
    <Sidebar />
    <MapSheet />
  </div>
);

export default App;
