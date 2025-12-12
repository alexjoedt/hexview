export namespace main {
	
	export class ConversionResult {
	    int8BE?: number;
	    int16BE?: number;
	    int32BE?: number;
	    int64BE?: number;
	    int16LE?: number;
	    int32LE?: number;
	    int64LE?: number;
	    uint8BE?: number;
	    uint16BE?: number;
	    uint32BE?: number;
	    uint64BE?: number;
	    uint16LE?: number;
	    uint32LE?: number;
	    uint64LE?: number;
	    float32BE?: string;
	    float64BE?: string;
	    float32LE?: string;
	    float64LE?: string;
	    binary?: string;
	    bytes?: string;
	
	    static createFrom(source: any = {}) {
	        return new ConversionResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.int8BE = source["int8BE"];
	        this.int16BE = source["int16BE"];
	        this.int32BE = source["int32BE"];
	        this.int64BE = source["int64BE"];
	        this.int16LE = source["int16LE"];
	        this.int32LE = source["int32LE"];
	        this.int64LE = source["int64LE"];
	        this.uint8BE = source["uint8BE"];
	        this.uint16BE = source["uint16BE"];
	        this.uint32BE = source["uint32BE"];
	        this.uint64BE = source["uint64BE"];
	        this.uint16LE = source["uint16LE"];
	        this.uint32LE = source["uint32LE"];
	        this.uint64LE = source["uint64LE"];
	        this.float32BE = source["float32BE"];
	        this.float64BE = source["float64BE"];
	        this.float32LE = source["float32LE"];
	        this.float64LE = source["float64LE"];
	        this.binary = source["binary"];
	        this.bytes = source["bytes"];
	    }
	}

}

