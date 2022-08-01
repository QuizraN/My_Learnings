import { useParams } from "react-router-dom";
import {useQuery} from "react-query"
import axios from "axios";
import '../App.css';

export const Products=()=>{
    const {id}=useParams()
    console.log(id)
    const { isLoading, isError, data, error }= useQuery(['products'],()=>{
        return axios.get(`https://obscure-refuge-62167.herokuapp.com/products/${id}`);
      })
      if (isLoading) {
        return <span>Loading...</span>
      }
    
      if (isError) {
        console.log(error)
        return <span>Error: {error.message}</span>
      }
      console.log(data.data)
    return (
        <div className="item">
            <img width={400} height={400} src={data.data.image} alt="product-1" />
            <div  className="description">
                <h4>Name:{data.data.name}</h4>
                <p>Description:{data.data.description}</p>
                <p>Price:{data.data.price}</p>
                <button id="btn" style={{marginTop:"30px",backgroundColor:"blue",color:"white",height:"30px",borderRadius:"4px"}}>Add to Cart</button>
            </div>   
        </div>
    );
}