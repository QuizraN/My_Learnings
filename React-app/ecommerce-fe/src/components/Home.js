import {useQuery} from "react-query"
import axios from "axios"
import {NavLink,useParams } from "react-router-dom";
import '../App.css';

export const Home=()=>{
    const {id}=useParams()
    const { isLoading, isError, data, error }= useQuery(['products'],()=>{
        return axios.get('https://obscure-refuge-62167.herokuapp.com/products');
      })
    if (isLoading) {
        return <span>Loading...</span>
    }
    
    if (isError) {
        console.log(error)
        return <span>Error: {error.message}</span>
    }
    console.log(data)
    return (
        <>
        
          <div className="nav2">
              <h1>Products</h1>
              <button style={{marginTop:"30px",backgroundColor:"blue",color:"white",height:"30px",borderRadius:"4px"}}>Create Product</button>
          </div>
          <div className="products-list">
          {
          data.data.map((i)=>{
            const path=`/products/${i.id}`
            return (<>
            <NavLink className="link" to={path}><img width={400} height={500} src={i.image} alt="product" /></NavLink>
            </>)})
          }
          </div>
        
        </>
      )
  
  
    }