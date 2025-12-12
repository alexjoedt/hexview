export namespace main {
	
	export class ConversionResult {
	    int8BE?: number;
	    int8BEHex?: string;
	    int16BE?: number;
	    int16BEHex?: string;
	    int32BE?: number;
	    int32BEHex?: string;
	    int64BE?: number;
	    int64BEHex?: string;
	    int16LE?: number;
	    int16LEHex?: string;
	    int32LE?: number;
	    int32LEHex?: string;
	    int64LE?: number;
	    int64LEHex?: string;
	    int16BADC?: number;
	    int16BADCHex?: string;
	    int32BADC?: number;
	    int32BADCHex?: string;
	    int64BADC?: number;
	    int64BADCHex?: string;
	    int16CDAB?: number;
	    int16CDABHex?: string;
	    int32CDAB?: number;
	    int32CDABHex?: string;
	    int64CDAB?: number;
	    int64CDABHex?: string;
	    uint8BE?: number;
	    uint8BEHex?: string;
	    uint16BE?: number;
	    uint16BEHex?: string;
	    uint32BE?: number;
	    uint32BEHex?: string;
	    uint64BE?: number;
	    uint64BEHex?: string;
	    uint16LE?: number;
	    uint16LEHex?: string;
	    uint32LE?: number;
	    uint32LEHex?: string;
	    uint64LE?: number;
	    uint64LEHex?: string;
	    uint16BADC?: number;
	    uint16BADCHex?: string;
	    uint32BADC?: number;
	    uint32BADCHex?: string;
	    uint64BADC?: number;
	    uint64BADCHex?: string;
	    uint16CDAB?: number;
	    uint16CDABHex?: string;
	    uint32CDAB?: number;
	    uint32CDABHex?: string;
	    uint64CDAB?: number;
	    uint64CDABHex?: string;
	    float32BE?: string;
	    float32BEHex?: string;
	    float64BE?: string;
	    float64BEHex?: string;
	    float32LE?: string;
	    float32LEHex?: string;
	    float64LE?: string;
	    float64LEHex?: string;
	    float32BADC?: string;
	    float32BADCHex?: string;
	    float64BADC?: string;
	    float64BADCHex?: string;
	    float32CDAB?: string;
	    float32CDABHex?: string;
	    float64CDAB?: string;
	    float64CDABHex?: string;
	    binary?: string;
	    bytes?: string;
	    ascii?: string;
	
	    static createFrom(source: any = {}) {
	        return new ConversionResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.int8BE = source["int8BE"];
	        this.int8BEHex = source["int8BEHex"];
	        this.int16BE = source["int16BE"];
	        this.int16BEHex = source["int16BEHex"];
	        this.int32BE = source["int32BE"];
	        this.int32BEHex = source["int32BEHex"];
	        this.int64BE = source["int64BE"];
	        this.int64BEHex = source["int64BEHex"];
	        this.int16LE = source["int16LE"];
	        this.int16LEHex = source["int16LEHex"];
	        this.int32LE = source["int32LE"];
	        this.int32LEHex = source["int32LEHex"];
	        this.int64LE = source["int64LE"];
	        this.int64LEHex = source["int64LEHex"];
	        this.int16BADC = source["int16BADC"];
	        this.int16BADCHex = source["int16BADCHex"];
	        this.int32BADC = source["int32BADC"];
	        this.int32BADCHex = source["int32BADCHex"];
	        this.int64BADC = source["int64BADC"];
	        this.int64BADCHex = source["int64BADCHex"];
	        this.int16CDAB = source["int16CDAB"];
	        this.int16CDABHex = source["int16CDABHex"];
	        this.int32CDAB = source["int32CDAB"];
	        this.int32CDABHex = source["int32CDABHex"];
	        this.int64CDAB = source["int64CDAB"];
	        this.int64CDABHex = source["int64CDABHex"];
	        this.uint8BE = source["uint8BE"];
	        this.uint8BEHex = source["uint8BEHex"];
	        this.uint16BE = source["uint16BE"];
	        this.uint16BEHex = source["uint16BEHex"];
	        this.uint32BE = source["uint32BE"];
	        this.uint32BEHex = source["uint32BEHex"];
	        this.uint64BE = source["uint64BE"];
	        this.uint64BEHex = source["uint64BEHex"];
	        this.uint16LE = source["uint16LE"];
	        this.uint16LEHex = source["uint16LEHex"];
	        this.uint32LE = source["uint32LE"];
	        this.uint32LEHex = source["uint32LEHex"];
	        this.uint64LE = source["uint64LE"];
	        this.uint64LEHex = source["uint64LEHex"];
	        this.uint16BADC = source["uint16BADC"];
	        this.uint16BADCHex = source["uint16BADCHex"];
	        this.uint32BADC = source["uint32BADC"];
	        this.uint32BADCHex = source["uint32BADCHex"];
	        this.uint64BADC = source["uint64BADC"];
	        this.uint64BADCHex = source["uint64BADCHex"];
	        this.uint16CDAB = source["uint16CDAB"];
	        this.uint16CDABHex = source["uint16CDABHex"];
	        this.uint32CDAB = source["uint32CDAB"];
	        this.uint32CDABHex = source["uint32CDABHex"];
	        this.uint64CDAB = source["uint64CDAB"];
	        this.uint64CDABHex = source["uint64CDABHex"];
	        this.float32BE = source["float32BE"];
	        this.float32BEHex = source["float32BEHex"];
	        this.float64BE = source["float64BE"];
	        this.float64BEHex = source["float64BEHex"];
	        this.float32LE = source["float32LE"];
	        this.float32LEHex = source["float32LEHex"];
	        this.float64LE = source["float64LE"];
	        this.float64LEHex = source["float64LEHex"];
	        this.float32BADC = source["float32BADC"];
	        this.float32BADCHex = source["float32BADCHex"];
	        this.float64BADC = source["float64BADC"];
	        this.float64BADCHex = source["float64BADCHex"];
	        this.float32CDAB = source["float32CDAB"];
	        this.float32CDABHex = source["float32CDABHex"];
	        this.float64CDAB = source["float64CDAB"];
	        this.float64CDABHex = source["float64CDABHex"];
	        this.binary = source["binary"];
	        this.bytes = source["bytes"];
	        this.ascii = source["ascii"];
	    }
	}
	export class ModbusCombined32 {
	    registerStart: number;
	    hex: string;
	    uint32BE: number;
	    uint32LE: number;
	    uint32BADC: number;
	    uint32CDAB: number;
	    int32BE: number;
	    int32LE: number;
	    int32BADC: number;
	    int32CDAB: number;
	    float32BE: string;
	    float32LE: string;
	    float32BADC: string;
	    float32CDAB: string;
	
	    static createFrom(source: any = {}) {
	        return new ModbusCombined32(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.registerStart = source["registerStart"];
	        this.hex = source["hex"];
	        this.uint32BE = source["uint32BE"];
	        this.uint32LE = source["uint32LE"];
	        this.uint32BADC = source["uint32BADC"];
	        this.uint32CDAB = source["uint32CDAB"];
	        this.int32BE = source["int32BE"];
	        this.int32LE = source["int32LE"];
	        this.int32BADC = source["int32BADC"];
	        this.int32CDAB = source["int32CDAB"];
	        this.float32BE = source["float32BE"];
	        this.float32LE = source["float32LE"];
	        this.float32BADC = source["float32BADC"];
	        this.float32CDAB = source["float32CDAB"];
	    }
	}
	export class ModbusCombined64 {
	    registerStart: number;
	    hex: string;
	    uint64BE: number;
	    uint64LE: number;
	    int64BE: number;
	    int64LE: number;
	    float64BE: string;
	    float64LE: string;
	
	    static createFrom(source: any = {}) {
	        return new ModbusCombined64(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.registerStart = source["registerStart"];
	        this.hex = source["hex"];
	        this.uint64BE = source["uint64BE"];
	        this.uint64LE = source["uint64LE"];
	        this.int64BE = source["int64BE"];
	        this.int64LE = source["int64LE"];
	        this.float64BE = source["float64BE"];
	        this.float64LE = source["float64LE"];
	    }
	}
	export class ModbusRegister {
	    index: number;
	    hex: string;
	    unsigned: number;
	    signed: number;
	    binary: string;
	
	    static createFrom(source: any = {}) {
	        return new ModbusRegister(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.index = source["index"];
	        this.hex = source["hex"];
	        this.unsigned = source["unsigned"];
	        this.signed = source["signed"];
	        this.binary = source["binary"];
	    }
	}
	export class ModbusResult {
	    registers: ModbusRegister[];
	    combined32: ModbusCombined32[];
	    combined64: ModbusCombined64[];
	    rawHex: string;
	    ascii: string;
	
	    static createFrom(source: any = {}) {
	        return new ModbusResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.registers = this.convertValues(source["registers"], ModbusRegister);
	        this.combined32 = this.convertValues(source["combined32"], ModbusCombined32);
	        this.combined64 = this.convertValues(source["combined64"], ModbusCombined64);
	        this.rawHex = source["rawHex"];
	        this.ascii = source["ascii"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

