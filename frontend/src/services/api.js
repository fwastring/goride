import axios from 'axios';

const BASE_URL = '/api';

// Create Axios instances for each API group
const userApi = axios.create({
  baseURL: `${BASE_URL}/user`,
  headers: {
    'Content-Type': 'application/json',
  }
});

const routeApi = axios.create({
  baseURL: `${BASE_URL}/route`,
  headers: {
    'Content-Type': 'application/json',
  }
});

const routesApi = axios.create({
  baseURL: `${BASE_URL}/routes`,
  headers: {
    'Content-Type': 'application/json',
  }
});

// Example User API functions
export const getUser = (id) => {
  return userApi.get(`/${id}`);
};

export const createUser = (userData) => {
  return userApi.post('/', userData);
};

export const updateUser = (id, userData) => {
  return userApi.put(`/${id}`, userData);
};

export const deleteUser = (id) => {
  return userApi.delete(`/${id}`);
};

// Example Route API functions
export const getRoute = (id) => {
  return routeApi.get(`/${id}`);
};

export const createRoute = (routeData) => {
  return routeApi.post('/', routeData);
};

export const updateRoute = (id, routeData) => {
  return routeApi.put(`/${id}`, routeData);
};

export const deleteRoute = (id) => {
  return routeApi.delete(`/${id}`);
};


// Example Routes API functions
export const getRoutes = () => {
  return routesApi.get();
};
