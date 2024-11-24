import { CheerioCrawler, CheerioCrawlingContext, RequestQueue } from "crawlee";

const callback = (err: Error, res, done) => {
	if (err) {
		console.log(err.stack);
		return done();
	}

	const $ = res.$;

	console.log(res);

	done();
};

export const Single = async (bookUrl: string) => {
	console.log(bookUrl);

	const requestQueue = await RequestQueue.open();
	await requestQueue.addRequest({ url: bookUrl });

	const crawler = new CheerioCrawler({
		requestQueue,
		additionalMimeTypes: ["application/xhtml+xml"],
		async requestHandler({ $, request }: CheerioCrawlingContext) {
			const titlePath = "head > meta[property='og:title']";
			const descriptionPath = "head > meta[property='og:description']";
			const coverPath = "head > meta[property='og:image']";
			const categoryPath = "head > meta[property='og:novel:category']";
			const authorPath = "head > meta[property='og:novel:author']";
			const statusPath = "head > meta[property='og:novel:status']";
			const lastChapterNamePath = "head > meta[property='og:novel:last_chapter_name']";
			const lastChapterUrlPath = "head > meta[property='og:novel:latest_chapter_url']";
			const lastUpdateTimePath = "head > meta[property='og:novel:update_time']";

			console.log(request);
			console.log("title:", $(titlePath));
			console.log("title1:", $(titlePath).attr("content"));
		},
	});

	await crawler.run();
};
