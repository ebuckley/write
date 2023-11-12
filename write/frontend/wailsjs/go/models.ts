export namespace main {
	
	export class FileWithContent {
	    name?: string;
	    content?: string;
	    html?: string;
	
	    static createFrom(source: any = {}) {
	        return new FileWithContent(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.content = source["content"];
	        this.html = source["html"];
	    }
	}

}

