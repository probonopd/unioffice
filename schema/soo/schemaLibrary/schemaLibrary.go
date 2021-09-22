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

package schemaLibrary ;import (_c "encoding/xml";_gb "fmt";_gg "github.com/unidoc/unioffice";);func (_b *CT_Schema )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for _ ,_cc :=range start .Attr {if _cc .Name .Local =="\u0075\u0072\u0069"{_e ,_d :=_cc .Value ,error (nil );if _d !=nil {return _d ;};_b .UriAttr =&_e ;continue ;};if _cc .Name .Local =="\u006d\u0061n\u0069\u0066\u0065s\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"{_f ,_ea :=_cc .Value ,error (nil );if _ea !=nil {return _ea ;};_b .ManifestLocationAttr =&_f ;continue ;};if _cc .Name .Local =="\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"{_df ,_fc :=_cc .Value ,error (nil );if _fc !=nil {return _fc ;};_b .SchemaLocationAttr =&_df ;continue ;};if _cc .Name .Local =="\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"{_ggf ,_ff :=_cc .Value ,error (nil );if _ff !=nil {return _ff ;};_b .SchemaLanguageAttr =&_ggf ;continue ;};};for {_ca ,_cd :=d .Token ();if _cd !=nil {return _gb .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0063\u0068e\u006d\u0061\u003a\u0020\u0025\u0073",_cd );};if _fb ,_ce :=_ca .(_c .EndElement );_ce &&_fb .Name ==start .Name {break ;};};return nil ;};func NewCT_SchemaLibrary ()*CT_SchemaLibrary {_da :=&CT_SchemaLibrary {};return _da };func (_a *CT_SchemaLibrary )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {e .EncodeToken (start );if _a .Schema !=nil {_eaa :=_c .StartElement {Name :_c .Name {Local :"\u006da\u003a\u0073\u0063\u0068\u0065\u006da"}};for _ ,_ba :=range _a .Schema {e .EncodeElement (_ba ,_eaa );};};e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};type SchemaLibrary struct{CT_SchemaLibrary };

// Validate validates the CT_Schema and its children
func (_bd *CT_Schema )Validate ()error {return _bd .ValidateWithPath ("\u0043T\u005f\u0053\u0063\u0068\u0065\u006da");};func NewSchemaLibrary ()*SchemaLibrary {_dag :=&SchemaLibrary {};_dag .CT_SchemaLibrary =*NewCT_SchemaLibrary ();return _dag ;};type CT_SchemaLibrary struct{Schema []*CT_Schema ;};func (_ceg *SchemaLibrary )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"});start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u006d\u0061"},Value :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"});start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u006d\u0061:\u0073\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079";return _ceg .CT_SchemaLibrary .MarshalXML (e ,start );};

// ValidateWithPath validates the CT_Schema and its children, prefixing error messages with path
func (_fcg *CT_Schema )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the CT_SchemaLibrary and its children, prefixing error messages with path
func (_dbf *CT_SchemaLibrary )ValidateWithPath (path string )error {for _ffg ,_agb :=range _dbf .Schema {if _fd :=_agb .ValidateWithPath (_gb .Sprintf ("\u0025\u0073\u002f\u0053\u0063\u0068\u0065\u006d\u0061\u005b\u0025\u0064\u005d",path ,_ffg ));_fd !=nil {return _fd ;};};return nil ;};type CT_Schema struct{UriAttr *string ;ManifestLocationAttr *string ;SchemaLocationAttr *string ;SchemaLanguageAttr *string ;};

// Validate validates the CT_SchemaLibrary and its children
func (_fa *CT_SchemaLibrary )Validate ()error {return _fa .ValidateWithPath ("\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079");};

// Validate validates the SchemaLibrary and its children
func (_ffgc *SchemaLibrary )Validate ()error {return _ffgc .ValidateWithPath ("\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079");};func (_ga *CT_Schema )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _ga .UriAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u006d\u0061\u003a\u0075\u0072\u0069"},Value :_gb .Sprintf ("\u0025\u0076",*_ga .UriAttr )});};if _ga .ManifestLocationAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u006d\u0061\u003a\u006dan\u0069\u0066\u0065\u0073\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"},Value :_gb .Sprintf ("\u0025\u0076",*_ga .ManifestLocationAttr )});};if _ga .SchemaLocationAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"},Value :_gb .Sprintf ("\u0025\u0076",*_ga .SchemaLocationAttr )});};if _ga .SchemaLanguageAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"},Value :_gb .Sprintf ("\u0025\u0076",*_ga .SchemaLanguageAttr )});};e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_eaf *SchemaLibrary )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_eaf .CT_SchemaLibrary =*NewCT_SchemaLibrary ();_cde :for {_dbb ,_bdbf :=d .Token ();if _bdbf !=nil {return _bdbf ;};switch _agbg :=_dbb .(type ){case _c .StartElement :switch _agbg .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e",Local :"\u0073\u0063\u0068\u0065\u006d\u0061"}:_eb :=NewCT_Schema ();if _ee :=d .DecodeElement (_eb ,&_agbg );_ee !=nil {return _ee ;};_eaf .Schema =append (_eaf .Schema ,_eb );default:_gg .Log ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0063\u0068\u0065m\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079 \u0025\u0076",_agbg .Name );if _fce :=d .Skip ();_fce !=nil {return _fce ;};};case _c .EndElement :break _cde ;case _c .CharData :};};return nil ;};func (_fg *CT_SchemaLibrary )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_gga :for {_bdb ,_cf :=d .Token ();if _cf !=nil {return _cf ;};switch _ag :=_bdb .(type ){case _c .StartElement :switch _ag .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e",Local :"\u0073\u0063\u0068\u0065\u006d\u0061"}:_cdb :=NewCT_Schema ();if _db :=d .DecodeElement (_cdb ,&_ag );_db !=nil {return _db ;};_fg .Schema =append (_fg .Schema ,_cdb );default:_gg .Log ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u0020\u0025v",_ag .Name );if _ad :=d .Skip ();_ad !=nil {return _ad ;};};case _c .EndElement :break _gga ;case _c .CharData :};};return nil ;};func NewCT_Schema ()*CT_Schema {_ge :=&CT_Schema {};return _ge };

// ValidateWithPath validates the SchemaLibrary and its children, prefixing error messages with path
func (_ae *SchemaLibrary )ValidateWithPath (path string )error {if _bc :=_ae .CT_SchemaLibrary .ValidateWithPath (path );_bc !=nil {return _bc ;};return nil ;};func init (){_gg .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0043T\u005f\u0053\u0063\u0068\u0065\u006da",NewCT_Schema );_gg .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079",NewCT_SchemaLibrary );_gg .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079",NewSchemaLibrary );};