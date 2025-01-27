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

package terms ;import (_ga "encoding/xml";_c "fmt";_e "github.com/unidoc/unioffice";_gc "github.com/unidoc/unioffice/common/logger";_d "github.com/unidoc/unioffice/schema/purl.org/dc/elements";);

// Validate validates the IMT and its children
func (_aca *IMT )Validate ()error {return _aca .ValidateWithPath ("\u0049\u004d\u0054")};type URI struct{};func (_dc *ElementOrRefinementContainer )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072";
e .EncodeToken (start );if _dc .Choice !=nil {for _ ,_bba :=range _dc .Choice {_bba .MarshalXML (e ,_ga .StartElement {});};};e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};

// Validate validates the Period and its children
func (_adbb *Period )Validate ()error {return _adbb .ValidateWithPath ("\u0050\u0065\u0072\u0069\u006f\u0064");};

// Validate validates the ElementsAndRefinementsGroup and its children
func (_fag *ElementsAndRefinementsGroup )Validate ()error {return _fag .ValidateWithPath ("E\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070");};func NewMESH ()*MESH {_eeb :=&MESH {};
return _eeb };

// Validate validates the ElementOrRefinementContainer and its children
func (_cf *ElementOrRefinementContainer )Validate ()error {return _cf .ValidateWithPath ("\u0045\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072");};func (_ag *ElementsAndRefinementsGroupChoice )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {_gbc :for {_eb ,_fee :=d .Token ();
if _fee !=nil {return _fee ;};switch _agc :=_eb .(type ){case _ga .StartElement :switch _agc .Name {case _ga .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_cca :=_d .NewAny ();
if _ca :=d .DecodeElement (_cca ,&_agc );_ca !=nil {return _ca ;};_ag .Any =append (_ag .Any ,_cca );default:_gc .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0041\u006ed\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006fu\u0070\u0043\u0068o\u0069\u0063\u0065\u0020\u0025\u0076",_agc .Name );
if _dea :=d .Skip ();_dea !=nil {return _dea ;};};case _ga .EndElement :break _gbc ;case _ga .CharData :};};return nil ;};

// ValidateWithPath validates the LCC and its children, prefixing error messages with path
func (_bg *LCC )ValidateWithPath (path string )error {return nil };func (_cd *Box )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_f ,_fd :=d .Token ();if _fd !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0042\u006f\u0078\u003a\u0020\u0025\u0073",_fd );
};if _ee ,_a :=_f .(_ga .EndElement );_a &&_ee .Name ==start .Name {break ;};};return nil ;};type IMT struct{};func (_deg *DCMIType )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065";
e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};func (_gab *DCMIType )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_b ,_def :=d .Token ();if _def !=nil {return _c .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0044\u0043\u004dI\u0054\u0079\u0070\u0065: \u0025\u0073",_def );
};if _ea ,_cb :=_b .(_ga .EndElement );_cb &&_ea .Name ==start .Name {break ;};};return nil ;};func (_cc *ElementsAndRefinementsGroupChoice )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {if _cc .Any !=nil {_cg :=_ga .StartElement {Name :_ga .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};
for _ ,_fgc :=range _cc .Any {e .EncodeElement (_fgc ,_cg );};};return nil ;};type LCSH struct{};

// ValidateWithPath validates the ElementOrRefinementContainer and its children, prefixing error messages with path
func (_ddd *ElementOrRefinementContainer )ValidateWithPath (path string )error {for _bfc ,_ff :=range _ddd .Choice {if _da :=_ff .ValidateWithPath (_c .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_bfc ));
_da !=nil {return _da ;};};return nil ;};

// ValidateWithPath validates the UDC and its children, prefixing error messages with path
func (_dae *UDC )ValidateWithPath (path string )error {return nil };func (_bgg *LCSH )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u004c\u0043\u0053\u0048";e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });
return nil ;};

// ValidateWithPath validates the Period and its children, prefixing error messages with path
func (_agd *Period )ValidateWithPath (path string )error {return nil };func NewDCMIType ()*DCMIType {_ac :=&DCMIType {};return _ac };

// ValidateWithPath validates the ElementsAndRefinementsGroupChoice and its children, prefixing error messages with path
func (_gad *ElementsAndRefinementsGroupChoice )ValidateWithPath (path string )error {for _fgcg ,_bbd :=range _gad .Any {if _ceg :=_bbd .ValidateWithPath (_c .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_fgcg ));_ceg !=nil {return _ceg ;
};};return nil ;};func NewUDC ()*UDC {_adbg :=&UDC {};return _adbg };func (_abg *RFC1766 )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0052F\u0043\u0031\u0037\u0036\u0036";e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });
return nil ;};

// ValidateWithPath validates the ElementsAndRefinementsGroup and its children, prefixing error messages with path
func (_dee *ElementsAndRefinementsGroup )ValidateWithPath (path string )error {for _eag ,_abb :=range _dee .Choice {if _aa :=_abb .ValidateWithPath (_c .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_eag ));
_aa !=nil {return _aa ;};};return nil ;};func NewDDC ()*DDC {_ge :=&DDC {};return _ge };func (_ccf *ISO3166 )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_bfe ,_bfbb :=d .Token ();if _bfbb !=nil {return _c .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0049\u0053\u004f\u0033\u0031\u0036\u0036\u003a\u0020\u0025\u0073",_bfbb );
};if _eac ,_gd :=_bfe .(_ga .EndElement );_gd &&_eac .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the RFC1766 and its children, prefixing error messages with path
func (_eaa *RFC1766 )ValidateWithPath (path string )error {return nil };func (_adc *UDC )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_acf ,_gcg :=d .Token ();if _gcg !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0044\u0043\u003a\u0020\u0025\u0073",_gcg );
};if _bdae ,_edf :=_acf .(_ga .EndElement );_edf &&_bdae .Name ==start .Name {break ;};};return nil ;};func (_ecb *TGN )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_afd ,_acd :=d .Token ();if _acd !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0054\u0047\u004e\u003a\u0020\u0025\u0073",_acd );
};if _bad ,_cbe :=_afd .(_ga .EndElement );_cbe &&_bad .Name ==start .Name {break ;};};return nil ;};type RFC3066 struct{};

// Validate validates the RFC1766 and its children
func (_fca *RFC1766 )Validate ()error {return _fca .ValidateWithPath ("\u0052F\u0043\u0031\u0037\u0036\u0036");};func NewURI ()*URI {_cdb :=&URI {};return _cdb };func (_feg *LCC )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u004c\u0043\u0043";
e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};func (_cab *Point )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0050\u006f\u0069n\u0074";e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });
return nil ;};type ElementOrRefinementContainer struct{Choice []*ElementsAndRefinementsGroupChoice ;};func NewElementsAndRefinementsGroupChoice ()*ElementsAndRefinementsGroupChoice {_dcf :=&ElementsAndRefinementsGroupChoice {};return _dcf ;};type MESH struct{};


// Validate validates the MESH and its children
func (_degb *MESH )Validate ()error {return _degb .ValidateWithPath ("\u004d\u0045\u0053\u0048")};

// ValidateWithPath validates the IMT and its children, prefixing error messages with path
func (_ecc *IMT )ValidateWithPath (path string )error {return nil };func (_ba *DDC )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_bc ,_gb :=d .Token ();if _gb !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0044\u0044\u0043\u003a\u0020\u0025\u0073",_gb );
};if _be ,_gea :=_bc .(_ga .EndElement );_gea &&_be .Name ==start .Name {break ;};};return nil ;};func (_age *IMT )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0049\u004d\u0054";e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });
return nil ;};func NewISO3166 ()*ISO3166 {_gfc :=&ISO3166 {};return _gfc };

// Validate validates the RFC3066 and its children
func (_dfc *RFC3066 )Validate ()error {return _dfc .ValidateWithPath ("\u0052F\u0043\u0033\u0030\u0036\u0036");};func (_egg *Period )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0050\u0065\u0072\u0069\u006f\u0064";
e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};type UDC struct{};func (_gbd *Period )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_feea ,_fbf :=d .Token ();if _fbf !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0050e\u0072\u0069o\u0064\u003a\u0020\u0025\u0073",_fbf );
};if _gcf ,_adf :=_feea .(_ga .EndElement );_adf &&_gcf .Name ==start .Name {break ;};};return nil ;};

// Validate validates the ElementsAndRefinementsGroupChoice and its children
func (_ada *ElementsAndRefinementsGroupChoice )Validate ()error {return _ada .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006et\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0043h\u006f\u0069\u0063\u0065");
};type TGN struct{};func (_ecd *URI )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_geff ,_fce :=d .Token ();if _fce !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0052\u0049\u003a\u0020\u0025\u0073",_fce );
};if _bca ,_cdg :=_geff .(_ga .EndElement );_cdg &&_bca .Name ==start .Name {break ;};};return nil ;};func NewElementsAndRefinementsGroup ()*ElementsAndRefinementsGroup {_bfcf :=&ElementsAndRefinementsGroup {};return _bfcf ;};type Point struct{};

// Validate validates the URI and its children
func (_aae *URI )Validate ()error {return _aae .ValidateWithPath ("\u0055\u0052\u0049")};func (_gff *MESH )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u004d\u0045\u0053\u0048";e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });
return nil ;};func (_gge *RFC1766 )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_acbf ,_fdaa :=d .Token ();if _fdaa !=nil {return _c .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0031\u0037\u0036\u0036\u003a\u0020\u0025\u0073",_fdaa );
};if _fba ,_egf :=_acbf .(_ga .EndElement );_egf &&_fba .Name ==start .Name {break ;};};return nil ;};func (_cfe *ISO3166 )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0049S\u004f\u0033\u0031\u0036\u0036";e .EncodeToken (start );
e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};func NewBox ()*Box {_ed :=&Box {};return _ed };

// ValidateWithPath validates the LCSH and its children, prefixing error messages with path
func (_egb *LCSH )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the DCMIType and its children, prefixing error messages with path
func (_bf *DCMIType )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the Box and its children, prefixing error messages with path
func (_dd *Box )ValidateWithPath (path string )error {return nil };func NewLCC ()*LCC {_bda :=&LCC {};return _bda };

// Validate validates the DDC and its children
func (_bb *DDC )Validate ()error {return _bb .ValidateWithPath ("\u0044\u0044\u0043")};func NewTGN ()*TGN {_egc :=&TGN {};return _egc };func NewElementOrRefinementContainer ()*ElementOrRefinementContainer {_ad :=&ElementOrRefinementContainer {};return _ad ;
};

// Validate validates the UDC and its children
func (_fcc *UDC )Validate ()error {return _fcc .ValidateWithPath ("\u0055\u0044\u0043")};

// Validate validates the ISO639_2 and its children
func (_gdc *ISO639_2 )Validate ()error {return _gdc .ValidateWithPath ("\u0049\u0053\u004f\u0036\u0033\u0039\u005f\u0032");};

// Validate validates the LCC and its children
func (_cbc *LCC )Validate ()error {return _cbc .ValidateWithPath ("\u004c\u0043\u0043")};func (_edde *RFC3066 )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_gec ,_afg :=d .Token ();if _afg !=nil {return _c .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0033\u0030\u0036\u0036\u003a\u0020\u0025\u0073",_afg );
};if _feeg ,_aec :=_gec .(_ga .EndElement );_aec &&_feeg .Name ==start .Name {break ;};};return nil ;};func (_dda *ISO639_2 )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032";
e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};func (_fae *Point )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_gee ,_eaf :=d .Token ();if _eaf !=nil {return _c .Errorf ("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0050\u006f\u0069\u006et\u003a\u0020\u0025\u0073",_eaf );
};if _fgb ,_ddfb :=_gee .(_ga .EndElement );_ddfb &&_fgb .Name ==start .Name {break ;};};return nil ;};func (_gfff *MESH )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_egbd ,_afb :=d .Token ();if _afb !=nil {return _c .Errorf ("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004d\u0045\u0053\u0048\u003a\u0020\u0025\u0073",_afb );
};if _fcb ,_add :=_egbd .(_ga .EndElement );_add &&_fcb .Name ==start .Name {break ;};};return nil ;};

// Validate validates the LCSH and its children
func (_gaf *LCSH )Validate ()error {return _gaf .ValidateWithPath ("\u004c\u0043\u0053\u0048")};

// ValidateWithPath validates the DDC and its children, prefixing error messages with path
func (_bee *DDC )ValidateWithPath (path string )error {return nil };

// Validate validates the W3CDTF and its children
func (_defag *W3CDTF )Validate ()error {return _defag .ValidateWithPath ("\u0057\u0033\u0043\u0044\u0054\u0046");};

// Validate validates the DCMIType and its children
func (_edg *DCMIType )Validate ()error {return _edg .ValidateWithPath ("\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065");};

// ValidateWithPath validates the Point and its children, prefixing error messages with path
func (_bdd *Point )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the TGN and its children, prefixing error messages with path
func (_fbe *TGN )ValidateWithPath (path string )error {return nil };type ISO639_2 struct{};

// Validate validates the Box and its children
func (_gg *Box )Validate ()error {return _gg .ValidateWithPath ("\u0042\u006f\u0078")};type Box struct{};type DDC struct{};type ElementsAndRefinementsGroupChoice struct{Any []*_d .Any ;};

// ValidateWithPath validates the URI and its children, prefixing error messages with path
func (_gaba *URI )ValidateWithPath (path string )error {return nil };func NewIMT ()*IMT {_ec :=&IMT {};return _ec };func NewRFC3066 ()*RFC3066 {_fbad :=&RFC3066 {};return _fbad };

// ValidateWithPath validates the ISO3166 and its children, prefixing error messages with path
func (_ebf *ISO3166 )ValidateWithPath (path string )error {return nil };func NewPeriod ()*Period {_db :=&Period {};return _db };type LCC struct{};func (_edgb *DDC )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0044\u0044\u0043";
e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the ISO639_2 and its children, prefixing error messages with path
func (_cbf *ISO639_2 )ValidateWithPath (path string )error {return nil };func NewPoint ()*Point {_fdg :=&Point {};return _fdg };func (_eaab *URI )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0055\u0052\u0049";e .EncodeToken (start );
e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};func (_bfb *IMT )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_fad ,_ddf :=d .Token ();if _ddf !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0049\u004d\u0054\u003a\u0020\u0025\u0073",_ddf );
};if _cfc ,_ebb :=_fad .(_ga .EndElement );_ebb &&_cfc .Name ==start .Name {break ;};};return nil ;};func (_bab *ElementsAndRefinementsGroup )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {if _bab .Choice !=nil {for _ ,_gbg :=range _bab .Choice {_gbg .MarshalXML (e ,_ga .StartElement {});
};};return nil ;};func (_gabc *ElementsAndRefinementsGroup )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {_eg :for {_fdc ,_dcc :=d .Token ();if _dcc !=nil {return _dcc ;};switch _df :=_fdc .(type ){case _ga .StartElement :switch _df .Name {case _ga .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_cff :=NewElementsAndRefinementsGroupChoice ();
if _cdc :=d .DecodeElement (&_cff .Any ,&_df );_cdc !=nil {return _cdc ;};_gabc .Choice =append (_gabc .Choice ,_cff );default:_gc .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006ce\u006d\u0065\u006e\u0074\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0020\u0025\u0076",_df .Name );
if _fe :=d .Skip ();_fe !=nil {return _fe ;};};case _ga .EndElement :break _eg ;case _ga .CharData :};};return nil ;};func (_gbcb *W3CDTF )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0057\u0033\u0043\u0044\u0054\u0046";
e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};func NewW3CDTF ()*W3CDTF {_bae :=&W3CDTF {};return _bae };func NewISO639_2 ()*ISO639_2 {_eef :=&ISO639_2 {};return _eef };func (_bbb *ISO639_2 )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_fb ,_fc :=d .Token ();
if _fc !=nil {return _c .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0049\u0053\u004f6\u0033\u0039\u005f\u0032: \u0025\u0073",_fc );};if _acb ,_adb :=_fb .(_ga .EndElement );_adb &&_acb .Name ==start .Name {break ;};};return nil ;};

// Validate validates the Point and its children
func (_gef *Point )Validate ()error {return _gef .ValidateWithPath ("\u0050\u006f\u0069n\u0074")};type ElementsAndRefinementsGroup struct{Choice []*ElementsAndRefinementsGroupChoice ;};type DCMIType struct{};func NewRFC1766 ()*RFC1766 {_ggb :=&RFC1766 {};
return _ggb };

// ValidateWithPath validates the RFC3066 and its children, prefixing error messages with path
func (_fegd *RFC3066 )ValidateWithPath (path string )error {return nil };func (_cgf *RFC3066 )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0052F\u0043\u0033\u0030\u0036\u0036";e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });
return nil ;};

// Validate validates the TGN and its children
func (_baba *TGN )Validate ()error {return _baba .ValidateWithPath ("\u0054\u0047\u004e")};func (_ceb *TGN )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0054\u0047\u004e";e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });
return nil ;};type RFC1766 struct{};type W3CDTF struct{};func (_de *Box )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0042\u006f\u0078";e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });
return nil ;};func (_ddg *LCC )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_ecce ,_edd :=d .Token ();if _edd !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u004c\u0043\u0043\u003a\u0020\u0025\u0073",_edd );};if _fda ,_gcb :=_ecce .(_ga .EndElement );
_gcb &&_fda .Name ==start .Name {break ;};};return nil ;};func (_dcd *W3CDTF )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_ccfa ,_acc :=d .Token ();if _acc !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u00573\u0043\u0044T\u0046\u003a\u0020\u0025\u0073",_acc );
};if _ef ,_cgd :=_ccfa .(_ga .EndElement );_cgd &&_ef .Name ==start .Name {break ;};};return nil ;};func NewLCSH ()*LCSH {_cdcb :=&LCSH {};return _cdcb };func (_bcg *UDC )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Name .Local ="\u0055\u0044\u0043";
e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the W3CDTF and its children, prefixing error messages with path
func (_cgfd *W3CDTF )ValidateWithPath (path string )error {return nil };func (_gfg *LCSH )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {for {_ae ,_defa :=d .Token ();if _defa !=nil {return _c .Errorf ("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004c\u0043\u0053\u0048\u003a\u0020\u0025\u0073",_defa );
};if _cfb ,_feb :=_ae .(_ga .EndElement );_feb &&_cfb .Name ==start .Name {break ;};};return nil ;};type ISO3166 struct{};func (_af *ElementOrRefinementContainer )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {_fa :for {_fg ,_dg :=d .Token ();
if _dg !=nil {return _dg ;};switch _fdb :=_fg .(type ){case _ga .StartElement :switch _fdb .Name {case _ga .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_gf :=NewElementsAndRefinementsGroupChoice ();
if _ce :=d .DecodeElement (&_gf .Any ,&_fdb );_ce !=nil {return _ce ;};_af .Choice =append (_af .Choice ,_gf );default:_gc .Log .Debug ("\u0073k\u0069\u0070\u0070\u0069\u006e\u0067\u0020un\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074 \u006f\u006e\u0020E\u006c\u0065\u006d\u0065\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065n\u0074\u0043on\u0074\u0061\u0069n\u0065\u0072\u0020\u0025\u0076",_fdb .Name );
if _adg :=d .Skip ();_adg !=nil {return _adg ;};};case _ga .EndElement :break _fa ;case _ga .CharData :};};return nil ;};

// Validate validates the ISO3166 and its children
func (_bd *ISO3166 )Validate ()error {return _bd .ValidateWithPath ("\u0049S\u004f\u0033\u0031\u0036\u0036");};type Period struct{};

// ValidateWithPath validates the MESH and its children, prefixing error messages with path
func (_deaf *MESH )ValidateWithPath (path string )error {return nil };func init (){_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004c\u0043\u0053\u0048",NewLCSH );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004d\u0045\u0053\u0048",NewMESH );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0044\u0044\u0043",NewDDC );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004c\u0043\u0043",NewLCC );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0055\u0044\u0043",NewUDC );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0050\u0065\u0072\u0069\u006f\u0064",NewPeriod );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0057\u0033\u0043\u0044\u0054\u0046",NewW3CDTF );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065",NewDCMIType );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049\u004d\u0054",NewIMT );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0055\u0052\u0049",NewURI );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032",NewISO639_2 );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0052F\u0043\u0031\u0037\u0036\u0036",NewRFC1766 );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0052F\u0043\u0033\u0030\u0036\u0036",NewRFC3066 );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0050\u006f\u0069n\u0074",NewPoint );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049S\u004f\u0033\u0031\u0036\u0036",NewISO3166 );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0042\u006f\u0078",NewBox );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0054\u0047\u004e",NewTGN );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072",NewElementOrRefinementContainer );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","e\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070",NewElementsAndRefinementsGroup );
};