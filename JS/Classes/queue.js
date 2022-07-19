class q
{
	enque(a,x)
	{
		a.push(x);
	}
	deque(a)
	{
		console.log(a.shift()+'is removed');
	}
	isFull(a,limit)
	{
		if(a.length==limit)
		{return true;}
		return false;
	}
	isEmpty(a)
	{
		if(a.length==0)
			return true;
		return false;
	}
}
	let a=[];
	let limit=6;
	var o1=new q();
	o1.enque(a,10);
	console.log(a);
	o1.deque(a);
	console.log(a);
	console.log(o1.isEmpty(a));
	console.log(o1.isFull(a,limit));


