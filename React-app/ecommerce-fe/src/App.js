import {Routes, Route, NavLink } from 'react-router-dom'
import {QueryClientProvider,QueryClient} from "react-query"

import {Home} from "./components/Home"
import {Products} from "./components/Products"
import {NotFound} from "./components/NotFound"
import './App.css';

const queryClient=new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
    <div className="main">
      <div className="nav1">
            <NavLink className={({isActive})=>( isActive ? "ActiveLink" : undefined) } to='/'>Home</NavLink>
            <NavLink className={({isActive})=>( isActive ? "ActiveLink" : undefined) } style={{marginLeft:"20px"}} to='/products/:id'>Products</NavLink>
      </div>  
      <Routes className="route">
            <Route exact path="/" element={<Home/>} />
            <Route path="/products/:id" element={<Products/>} />
            <Route exact path="*" element={<NotFound/>} />
      </Routes>
    </div>
    </QueryClientProvider>
  );
}
export default App;
