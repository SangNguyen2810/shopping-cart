import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import CartPage from '@/pages/CartPage/CartPage';
import ErrorPage from '@/pages/ErrorPage/ErrorPage';
import './app.css';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<CartPage />} />
        <Route path="*" element={<ErrorPage />} />
      </Routes>
    </Router>
  );
}

export default App;
