package main

import "testing"

var vecs = [...]uint64{
	0xea1131cf6f3f5848,
	0x5dce1178cbcfd871,
	0x7909b3317143cca1,
	0x12e6e7c6690629b0,
	0x5db52d6956d79454,
	0x41d650fe26c36227,
	0x776f59a87da7fa56,
	0xa40cd9a0427d4c3b,
	0x737bb7e392aeb3f8,
	0xd96a856ff8939e34,
	0xbb4afede1a8974e7,
	0x5894ed72e5ec7a96,
	0x35634585bfa2aef9,
	0x9e71ceeb88a4d1b0,
	0x4d2160646528dd52,
	0xe8330f398042b526,
	0x575b4b5c80e71151,
	0x85de83dfc7e27619,
	0xc9b5f383cefa7071,
	0x86aad28a5c47b43c,
	0x75403f14caa6bc23,
	0xe5426cc5f60085d0,
	0x97d95f941b573318,
	0xd19138a42188fdc8,
	0xd5b67b8fc48d451c,
	0x72c34791017d26cd,
	0x3c3ef56bff5c4316,
	0xf04a149c61b2d433,
	0x80acf2553f9bd6ef,
	0xd39eb1308956f0f3,
	0xac2c8bce566c41b2,
	0x73f9777a894cc8de,
	0xc25883dfbc3e8c32,
	0x7607a7bb23480f64,
	0x1970e219fe6f3a1c,
	0x081b8556551c3d30,
	0xc4fcca3a453040ae,
	0x5b78f06984ccd8f3,
	0x18aef6100a436e3a,
	0xa078e0bc5c5a1cf9,
	0x99ec970bbcc8a053,
	0xc1dca1687cdb7d77,
	0x91fa2c92ff32021c,
	0x4898d13e444a76b5,
	0x06ced1c3fbd8b164,
	0x97c184ea558c457f,
	0x6172386913b18ed2,
	0x4a5ce0b499fac854,
	0x37dc9799e168febc,
	0xd814fd9e8fbe3bd5,
	0x2b475321ad3c3ba3,
	0x1fe5ba853d2d7184,
	0x58099374c2a02018,
	0xba0f271ea5bbeefc,
	0x19aa09073af0cc4d,
	0x4314609d598697b6,
	0x6bc8b6ff93f4303b,
	0xbd83f821302de212,
	0x5d287bbd0df58875,
	0x4f6931cab8286b9b,
	0xf0ee6db506adfcc4,
	0xc20c86e5796b6ec7,
	0xc9b9cd6bc5f90c5f,
	0x930c787cdabd59ae,
	0x2a4b01429b28e9d9,
	0x2ffa72b11c8f1b00,
	0xa3a099f21995cf43,
	0xecc4dfc2ccef7d23,
	0xc1c35b9195ac888e,
	0xe258bde65e0f2b8b,
	0x40f754e341b56faf,
	0x0945cf4d17e92a3b,
	0x2fd5690c42607829,
	0xe8f72455004f88a1,
	0x0bec70028f4351a3,
	0x963d3052695f0683,
	0x4a3f4678ad5bd368,
	0xa056190ce114941a,
	0xdd4ec8462f0e62aa,
	0xc203169ef467712b,
	0x418e72c5136b153d,
	0x1e647e8519e34b98,
	0xdbd952d065e70de9,
	0x5ce123aa2bc15ade,
	0x71c5d5b652d8de74,
	0x0ed73e0ffecd25d3,
	0x40df92f28837d7a3,
	0x13b8873d0dc41b67,
	0x611705753e89c200,
	0x575cad48048e0fa3,
	0x6498d8f531c8de0a,
	0xc177be58d4cc79dd,
	0xa647b6a1d6863d2b,
	0xf103fe26fc0f3296,
	0xbbb0a913c67ca7fa,
	0x44bd03e1b5ae530c,
	0xbcba5963214a46cf,
	0x61b3fced4ec354b5,
	0x682b56c9f03319de,
	0x8b8da63ae2436431,
	0x8d3c830437847db9,
	0x4a74e1ac665cee1b,
	0x61c35a076c4ec680,
	0x92b90dd44efd797c,
	0x6a3d7075f0f2d2d6,
	0xc72917016f50ad37,
	0xb2bb1abffdf8e463,
	0xd33e7f77f099521b,
	0x68cce04666f4d8e4,
	0xb9e379a014ed2cfb,
	0xcccc99a63aedf234,
	0x42e238364e01c9ef,
	0x4fce3a62ecc9e468,
	0xcf6a029e721a4483,
	0x00df6c6566001222,
	0x953e3acadedf83c7,
	0x7583c30fd341f9cf,
	0xed5a859c20017728,
	0x12088fa98e8db6c7,
	0x140c122022ce0e8e,
	0x9e1aac2f598f405d,
	0xe777cb6884aaba74,
	0x8862519f1753bd7d,
	0x7ca4afd5746e6bc8,
	0x538e19af125490dd,
	0xb4dd969941e7716a,
	0x6d7c9af2f43a1a73,
	0x01223b7d408e4807,
}

func TestT1ha(t *testing.T) {
	var data []byte
	tt := &TH1A{0x0102030405060708}
	for i, want := range vecs {
		data = append(data, byte(i))
		got := tt.Sum64(data[:i])
		if got != want {
			t.Errorf("Sum64(...%d)=%x, want %x\n", i, got, want)
		}
	}
}
