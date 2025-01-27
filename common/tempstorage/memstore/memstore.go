//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package memstore implements tempStorage interface
// by using memory as a storage
package memstore ;import (_d "encoding/hex";_bg "errors";_c "fmt";_cg "github.com/unidoc/unioffice/common/tempstorage";_b "io";_eg "io/ioutil";_ef "math/rand";_e "sync";);

// Read reads from the underlying memDataCell in order to implement Reader interface
func (_cb *memFile )Read (p []byte )(int ,error ){_f :=_cb ._ea ;_bd :=_cb ._ga ._ca ;_eae :=int64 (len (p ));if _eae > _bd {_eae =_bd ;p =p [:_eae ];};if _f >=_bd {return 0,_b .EOF ;};_ff :=_f +_eae ;if _ff >=_bd {_ff =_bd ;};_gc :=copy (p ,_cb ._ga ._aea [_f :_ff ]);
_cb ._ea =_ff ;return _gc ,nil ;};

// ReadAt reads from the underlying memDataCell at an offset provided in order to implement ReaderAt interface.
// It does not affect f.readOffset.
func (_df *memFile )ReadAt (p []byte ,readOffset int64 )(int ,error ){_ec :=_df ._ga ._ca ;_db :=int64 (len (p ));if _db > _ec {_db =_ec ;p =p [:_db ];};if readOffset >=_ec {return 0,_b .EOF ;};_a :=readOffset +_db ;if _a >=_ec {_a =_ec ;};_fc :=copy (p ,_df ._ga ._aea [readOffset :_a ]);
return _fc ,nil ;};

// Add reads a file from a disk and adds it to the storage
func (_efd *memStorage )Add (path string )error {_ ,_dgc :=_efd ._da .Load (path );if _dgc {return nil ;};_de ,_af :=_eg .ReadFile (path );if _af !=nil {return _af ;};_efd ._da .Store (path ,&memDataCell {_gg :path ,_aea :_de ,_ca :int64 (len (_de ))});
return nil ;};

// Open returns tempstorage File object by name
func (_dd *memStorage )Open (path string )(_cg .File ,error ){_bga ,_cc :=_dd ._da .Load (path );if !_cc {return nil ,_bg .New (_c .Sprintf ("\u0043\u0061\u006eno\u0074\u0020\u006f\u0070\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073",path ));
};return &memFile {_ga :_bga .(*memDataCell )},nil ;};func _aab (_ecd int )(string ,error ){_cf :=make ([]byte ,_ecd );if _ ,_fa :=_ef .Read (_cf );_fa !=nil {return "",_fa ;};return _d .EncodeToString (_cf ),nil ;};

// Close is not applicable in this implementation
func (_fb *memFile )Close ()error {return nil };type memFile struct{_ga *memDataCell ;_ea int64 ;};type memDataCell struct{_gg string ;_aea []byte ;_ca int64 ;};

// SetAsStorage sets temp storage as a memory storage
func SetAsStorage (){_bc :=memStorage {_da :_e .Map {}};_cg .SetAsStorage (&_bc )};

// TempFile creates a new empty file in the storage and returns it
func (_dg *memStorage )TempFile (dir ,pattern string )(_cg .File ,error ){_ce :=dir +"\u002f"+_ffd (pattern );_ed :=&memDataCell {_gg :_ce ,_aea :[]byte {}};_be :=&memFile {_ga :_ed };_dg ._da .Store (_ce ,_ed );return _be ,nil ;};type memStorage struct{_da _e .Map };
func _ffd (_dga string )string {_cd ,_ :=_aab (6);return _dga +_cd };

// RemoveAll removes all files according to the dir argument prefix
func (_bgae *memStorage )RemoveAll (dir string )error {_bgae ._da .Range (func (_ggg ,_ccg interface{})bool {_bgae ._da .Delete (_ggg );return true });return nil ;};

// TempDir creates a name for a new temp directory using a pattern argument
func (_dda *memStorage )TempDir (pattern string )(string ,error ){return _ffd (pattern ),nil };

// Name returns the filename of the underlying memDataCell
func (_gac *memFile )Name ()string {return _gac ._ga ._gg };

// Write writes to the end of the underlying memDataCell in order to implement Writer interface
func (_ae *memFile )Write (p []byte )(int ,error ){_ae ._ga ._aea =append (_ae ._ga ._aea ,p ...);_ae ._ga ._ca +=int64 (len (p ));return len (p ),nil ;};