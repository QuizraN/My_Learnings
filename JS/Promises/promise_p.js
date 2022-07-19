
function lowerCase(str)
{
	return new Promise((resolve,reject)=>{
		setTimeout(()=>{
			if(str!='')
			{resolve(str.toLowerCase())}
			else
			{reject('string is empty');}
		},5000);
	});
}
const promise=lowerCase('BEAUTIFUL CODE');
promise.then((str)=>{return str;}).then((str)=>console.log(str+"-"+str.length))
	.catch((str)=>console.log(str))
