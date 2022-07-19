async function getProcessedData(url)
{
	try
	{
	const p1= await downloadData(url);
	const p2= await processDataInWorker(p1);
	}
	catch(e)
	{
		await downloadFallbackData(url);
	}
}
