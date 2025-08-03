import React from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import Upload from './pages/Upload';
import { useSelector } from 'react-redux';
import type { RootState } from './app/store';
import Navbar from './components/Navbar';


const App = () => {
  const { token, user } = useSelector((state: RootState) => state.auth);

  // Route wrapper for auth + role-based access
  const PrivateRoute = ({
    children,
    roles,
  }: {
    children: React.ReactNode;
    roles?: string[];
  }) => {
    if (!token || !user) {
      return <Navigate to="/login" replace />;
    }

    if (roles && !roles.includes(user.role)) {
      return <Navigate to="/unauthorized" replace />;
    }

    return <>{children}</>;
  };

  return (
    <Routes>
      {/* Public */}
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
      <Route path="/unauthorized" element={<h2>403 - Unauthorized</h2>} />

      {/* Protected */}
      <Route
        path="/dashboard"
        element={
          <PrivateRoute>
            <Dashboard />
          </PrivateRoute>
        }
      />
      <Route
        path="/upload"
        element={
          <PrivateRoute roles={['admin', 'editor']}>
            <Upload />
          </PrivateRoute>
        }
      />

      {/* Catch-all redirect */}
      <Route path="*" element={<Navigate to="/dashboard" />} />
    </Routes>
  );
};

export default App;
