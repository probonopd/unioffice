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

package rw ;import (_e "bytes";_d "encoding/binary";_f "errors";_a "fmt";_gd "io";_ea "io/ioutil";_ga "reflect";);func (_dg *Reader )ReadStringProperty (n uint32 )(string ,error ){if _de :=_dg .align (4);_de !=nil {return "",_de ;};_bde :=make ([]byte ,n );
if _eabe :=_d .Read (_dg ,_d .LittleEndian ,&_bde );_eabe !=nil {return "",_eabe ;};return string (_bde ),nil ;};func (_gge *Writer )FillWithByte (fillSize int ,b byte )error {for _bec :=0;_bec < fillSize ;_bec ++{if _cb :=_gge .WritePropertyNoAlign (b );
_cb !=nil {return _cb ;};};return nil ;};type Reader struct{*_e .Reader };func NewWriter ()*Writer {return &Writer {_ed :[]byte {}}};func (_cc *Writer )WriteProperty (a interface{})error {if _eb :=_cc .align (int (_ga .TypeOf (a ).Size ()));_eb !=nil {return _eb ;
};return _cc .WritePropertyNoAlign (a );};func (_efb *Writer )curPos ()int {return int (_efb .Cap ())-_efb .Len ()};func (_fb *Writer )AlignLength (alignTo int )error {_deg :=_fb .Len ()%alignTo ;if _deg > 0{_ ,_bbg :=_fb .Write (make ([]byte ,alignTo -_deg ));
if _bbg !=nil {return _bbg ;};};return nil ;};func (_bbce *Writer )Skip (n int )error {if n ==0{return nil ;};_ ,_ce :=_bbce .Write (make ([]byte ,n ));return _ce ;};func (_gaa *Reader )curPos ()int {return int (_gaa .Size ())-_gaa .Len ()};func PopRightUI64 (v uint64 )(bool ,uint64 ){return (v &uint64 (1))==1,v >>1};
const _gdfb =64;func (_ab *Writer )WriteByteAt (b byte ,off int )error {if off >=len (_ab ._ed ){return _f .New ("\u004f\u0075\u0074\u0020\u006f\u0066\u0020\u0062\u006f\u0075\u006e\u0064\u0073");};_ab ._ed [off ]=b ;return nil ;};func (_fgg *Writer )WriteTo (wTo _gd .Writer )(_bfb int64 ,_ff error ){if _ag :=_fgg .Len ();
_ag > 0{_bfe ,_ccg :=wTo .Write (_fgg ._ed [_fgg ._fca :]);if _bfe > _ag {return 0,_f .New ("\u0072\u0077\u002e\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f\u003a\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0057\u0072\u0069t\u0065\u0020c\u006f\u0075\u006e\u0074");
};_fgg ._fca +=_bfe ;_bfb =int64 (_bfe );if _ccg !=nil {return _bfb ,_ccg ;};if _bfe !=_ag {return _bfb ,_gd .ErrShortWrite ;};};_fgg .reset ();return _bfb ,nil ;};var _bg =_f .New ("r\u0077.\u0057\u0072\u0069\u0074\u0065\u0072\u003a\u0020t\u006f\u006f\u0020\u006car\u0067\u0065");
func (_da *Writer )reset (){_da ._ed =_da ._ed [:0];_da ._fca =0};func (_fbc *Writer )Cap ()int {return cap (_fbc ._ed )};type Writer struct{_ed []byte ;_fca int ;};func _fge (_fbbb int )[]byte {defer func (){if recover ()!=nil {panic (_bg );};}();return make ([]byte ,_fbbb );
};func (_db *Reader )align (_df int )error {return _db .skip ((_df -_db .curPos ()%_df )%_df )};func (_bda *Writer )tryGrowByReslice (_cd int )(int ,bool ){if _ge :=len (_bda ._ed );_cd <=cap (_bda ._ed )-_ge {_bda ._ed =_bda ._ed [:_ge +_cd ];return _ge ,true ;
};return 0,false ;};func (_eab *Reader )ReadProperty (a interface{})error {_bb :=_ga .ValueOf (a );for _bb .Kind ()==_ga .Ptr {_bb =_bb .Elem ();};if !_bb .IsValid (){return _a .Errorf ("\u0076a\u006cu\u0065\u0020\u0069\u0073\u0020n\u006f\u0074 \u0076\u0061\u006c\u0069\u0064");
};if _gdf :=_eab .align (int (_bb .Type ().Size ()));_gdf !=nil {return _gdf ;};if _ef :=_d .Read (_eab ,_d .LittleEndian ,a );_ef !=nil {return _ef ;};return nil ;};func (_bc *Writer )Write (p []byte )(_fba int ,_fbb error ){_dc ,_gab :=_bc .tryGrowByReslice (len (p ));
if !_gab {var _dce error ;_dc ,_dce =_bc .grow (len (p ));if _dce !=nil {return 0,_dce ;};};return copy (_bc ._ed [_dc :],p ),nil ;};func PushLeftUI32 (v uint32 ,flag bool )uint32 {v >>=1;if flag {v |=1<<31;};return v ;};func NewReader (b []byte )(*Reader ,error ){return &Reader {_e .NewReader (b )},nil };
func (_be *Writer )align (_ccb int )error {return _be .Skip ((_ccb -(_be .Len ())%_ccb )%_ccb )};func (_gg *Writer )Len ()int {return len (_gg ._ed )-_gg ._fca };func PushLeftUI64 (v uint64 ,flag bool )uint64 {v >>=1;if flag {v |=1<<63;};return v ;};func (_fc *Reader )skip (_fg int )error {_ ,_b :=_gd .CopyN (_ea .Discard ,_fc ,int64 (_fg ));
if _b !=nil {return _b ;};return nil ;};func PopRightUI32 (v uint32 )(bool ,uint32 ){return (v &uint32 (1))==1,v >>1};func (_cda *Writer )grow (_fbd int )(int ,error ){_gb :=_cda .Len ();if _gb ==0&&_cda ._fca !=0{_cda .reset ();};if _ffb ,_fgf :=_cda .tryGrowByReslice (_fbd );
_fgf {return _ffb ,nil ;};if _cda ._ed ==nil &&_fbd <=_gdfb {_cda ._ed =make ([]byte ,_fbd ,_gdfb );return 0,nil ;};_cg :=cap (_cda ._ed );if _fbd <=_cg /2-_gb {copy (_cda ._ed ,_cda ._ed [_cda ._fca :]);}else if _cg > _dd -_cg -_fbd {return 0,_bg ;}else {_bca :=_fge (2*_cg +_fbd );
copy (_bca ,_cda ._ed [_cda ._fca :]);_cda ._ed =_bca ;};_cda ._fca =0;_cda ._ed =_cda ._ed [:_gb +_fbd ];return _gb ,nil ;};const _dd =int (^uint (0)>>1);func (_aa *Writer )WritePropertyNoAlign (a interface{})error {if _bbc :=_d .Write (_aa ,_d .LittleEndian ,a );
_bbc !=nil {return _bbc ;};return nil ;};func (_bbe *Writer )Bytes ()[]byte {return _bbe ._ed };func (_acc *Writer )WriteStringProperty (s string )error {_acc .align (4);_gc :=[]byte (s );if _bff :=_d .Write (_acc ,_d .LittleEndian ,&_gc );_bff !=nil {return _bff ;
};return nil ;};func (_bbf *Reader )ReadPairProperty (p interface{})error {if _bd :=_bbf .align (4);_bd !=nil {return _bd ;};_ac :=_ga .ValueOf (p );for _ac .Kind ()==_ga .Ptr {_ac =_ac .Elem ();};if !_ac .IsValid (){return _a .Errorf ("\u0076a\u006cu\u0065\u0020\u0069\u0073\u0020n\u006f\u0074 \u0076\u0061\u006c\u0069\u0064");
};if _bf :=_d .Read (_bbf ,_d .LittleEndian ,p );_bf !=nil {return _bf ;};return nil ;};