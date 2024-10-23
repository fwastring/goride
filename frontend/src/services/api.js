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
  return routeApi.post('/create', routeData);
};

export const searchRoute = (routeData) => {
  return routeApi.post('/search', routeData);
};

export const joinRoute = (routeData) => {
  return routeApi.post('/join', routeData);
};

export const updateRoute = (id, routeData) => {
  return routeApi.put(`/${id}`, routeData);
};

export const deleteRoute = (routeData) => {
  return routeApi.post(`/delete`, routeData);
};


// Example Routes API functions
export const getRoutes = () => {
  return routeApi.get(`/all`);
};
