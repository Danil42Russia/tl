// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"github.com/vkcom/tl/internal/utils"
	"golang.org/x/exp/slices"
	"strings"
)

func (trw *TypeRWStruct) CPPTypeJSONEmptyCondition(bytesVersion bool, val string, ref bool, deps []string) string {
	if trw.isTypeDef() {
		return trw.Fields[0].t.trw.CPPTypeJSONEmptyCondition(bytesVersion, val, ref, deps)
	}
	return ""
}

func (trw *TypeRWStruct) CPPFillRecursiveChildren(visitedNodes map[*TypeRWWrapper]bool) {
	for _, f := range trw.Fields {
		if !f.recursive {
			f.t.CPPFillRecursiveChildren(visitedNodes)
		}
	}
}

func (trw *TypeRWStruct) CPPAllowCurrentDefinition() bool {
	// todo
	if trw.isTypeDef() {
		typeReduction := trw.cppGetCurrentTypeReduction()

		field := trw.Fields[0]
		typeRed := trw.wr.gen.typesInfo.FieldTypeReduction(&typeReduction, 0)

		brackets, isBrackets := field.t.trw.(*TypeRWBrackets)
		if isBrackets {
			if brackets.dictLike && typeRed.Type.Arguments[0].Type != nil && len(typeRed.Type.Arguments[0].Type.Arguments) >= 1 {
				return false
			}
		}
	}
	return true
}

func (trw *TypeRWStruct) cppGetCurrentTypeReduction() TypeReduction {
	ti := trw.wr.gen.typesInfo
	tlName := trw.wr.tlName

	_, isType := ti.Types[tlName]
	typeReduction := TypeReduction{IsType: isType}
	if isType {
		typeReduction.Type = ti.Types[tlName]
	} else {
		typeReduction.Constructor = ti.Constructors[tlName]
	}
	for i, arg := range typeReduction.ReferenceType().TypeArguments {
		evalArg := EvaluatedType{}
		if arg.IsNat {
			evalArg.Index = 1
			evalArg.Variable = arg.FieldName
			if trw.wr.arguments[i].isArith {
				// set true only here
				evalArg.VariableActsAsConstant = true
			}
		} else {
			evalArg.Index = 3
			evalArg.TypeVariable = arg.FieldName
		}
		typeReduction.Arguments = append(typeReduction.Arguments, evalArg)
	}
	return typeReduction
}

func (trw *TypeRWStruct) cppTypeStringInNamespace(bytesVersion bool, hppInc *DirectIncludesCPP) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.CPPTypeStringInNamespace(bytesVersion, hppInc)
	}
	_, _, args := trw.wr.cppTypeStringInNamespace(bytesVersion, hppInc, false, HalfResolvedArgument{})
	return trw.wr.cppNamespaceQualifier() + trw.wr.cppLocalName + args
}

func (trw *TypeRWStruct) cppTypeStringInNamespaceHalfResolved(bytesVersion bool, hppInc *DirectIncludesCPP, halfResolved HalfResolvedArgument) string {
	if trw.isUnwrapType() {
		halfResolvedUnwrapped := trw.wr.replaceUnwrapHalfResolved(halfResolved, trw.Fields[0].halfResolved)
		return trw.Fields[0].t.CPPTypeStringInNamespaceHalfResolved(bytesVersion, hppInc, halfResolvedUnwrapped)
	}
	_, _, args := trw.wr.cppTypeStringInNamespace(bytesVersion, hppInc, true, halfResolved)
	return trw.wr.cppNamespaceQualifier() + trw.wr.cppLocalName + args
}

func (trw *TypeRWStruct) cppTypeStringInNamespaceHalfResolved2(bytesVersion bool, typeReduction EvaluatedType) string {
	if trw.isUnwrapType() {
		eval := trw.wr.gen.typesInfo.FieldTypeReduction(typeReduction.Type, 0)
		return trw.Fields[0].t.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, eval)
	}
	return trw.wr.cppNamespaceQualifier() + trw.wr.cppLocalName + trw.wr.cppTypeArguments(bytesVersion, typeReduction.Type)
}

func (trw *TypeRWStruct) cppDefaultInitializer(halfResolved HalfResolvedArgument, halfResolve bool) string {
	if trw.isUnwrapType() {
		return trw.Fields[0].t.CPPDefaultInitializer(trw.Fields[0].halfResolved, halfResolve)
	}
	return "{}"
}

func (trw *TypeRWStruct) CPPHasBytesVersion() bool {
	return false // TODO
}

func (trw *TypeRWStruct) CPPTypeResettingCode(bytesVersion bool, val string) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	if trw.isUnwrapType() {
		return trw.Fields[0].t.trw.CPPTypeResettingCode(bytesVersion, val)
	}
	return fmt.Sprintf("\t::%s::%sReset(%s);", trw.wr.gen.DetailsCPPNamespace, goGlobalName, val)
}

func (trw *TypeRWStruct) CPPTypeWritingJsonCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	if trw.isUnwrapType() {
		return trw.Fields[0].t.trw.CPPTypeWritingJsonCode(bytesVersion, val, trw.Fields[0].Bare(), trw.replaceUnwrapArgs(natArgs), last)
	}
	return fmt.Sprintf("\tif (!::%s::%sWriteJSON(s, %s%s)) { return false; }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, val, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) CPPTypeWritingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	if trw.isUnwrapType() {
		prefix := ""
		if !bare {
			prefix = fmt.Sprintf("\tif (!s.nat_write(0x%08x)) { return false; }\n", trw.wr.tlTag)
		}
		return prefix + trw.Fields[0].t.trw.CPPTypeWritingCode(bytesVersion, val, trw.Fields[0].Bare(), trw.replaceUnwrapArgs(natArgs), last)
	}
	return fmt.Sprintf("\tif (!::%s::%sWrite%s(s, %s%s)) { return s.set_error_unknown_scenario(); }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) CPPTypeReadingCode(bytesVersion bool, val string, bare bool, natArgs []string, last bool) string {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)
	if trw.isUnwrapType() {
		prefix := ""
		if !bare {
			prefix = fmt.Sprintf("\tif (!s.nat_read_exact_tag(0x%08x)) { return false; }\n", trw.wr.tlTag)
		}
		return prefix + trw.Fields[0].t.trw.CPPTypeReadingCode(bytesVersion, val, trw.Fields[0].Bare(), trw.replaceUnwrapArgs(natArgs), last)
	}
	return fmt.Sprintf("\tif (!::%s::%sRead%s(s, %s%s)) { return s.set_error_unknown_scenario(); }", trw.wr.gen.DetailsCPPNamespace, goGlobalName, addBare(bare), val, joinWithCommas(natArgs))
}

func (trw *TypeRWStruct) CPPGenerateCode(hpp *strings.Builder, hppInc *DirectIncludesCPP, hppIncFwd *DirectIncludesCPP, hppDet *strings.Builder, hppDetInc *DirectIncludesCPP, cppDet *strings.Builder, cppDetInc *DirectIncludesCPP, bytesVersion bool, forwardDeclaration bool) {
	goGlobalName := addBytes(trw.wr.goGlobalName, bytesVersion)

	_, myArgsDecl := trw.wr.fullyResolvedClassCppNameArgs()

	anyRecursive := false
	typeNamespace := trw.wr.gen.RootCPPNamespaceElements
	if trw.wr.tlName.Namespace != "" {
		typeNamespace = append(typeNamespace, trw.wr.tlName.Namespace)
	}
	if hpp != nil {
		if !forwardDeclaration {
			deps := trw.AllTypeDependencies(true, false)
			slices.SortFunc(deps, TypeComparator)

			for _, typeDep := range deps {
				if typeDep.typeComponent == trw.wr.typeComponent {
					typeDep.trw.CPPGenerateCode(hpp, nil, nil, nil, hppDetInc, nil, cppDetInc, bytesVersion, true)
				}
			}
		}
		cppStartNamespace(hpp, typeNamespace)
		// hpp.WriteString("// " + goLocalName + "\n") - uncommenting will lead to multiple definition error
		if len(myArgsDecl) != 0 {
			hpp.WriteString("template<" + strings.Join(myArgsDecl, ", ") + ">\n")
		}
		if forwardDeclaration { // TODO - does not work for typedef
			hpp.WriteString("struct " + trw.wr.cppLocalName + ";")
			cppFinishNamespace(hpp, typeNamespace)
			return
		}

		ti := trw.wr.gen.typesInfo
		typeReduction := trw.cppGetCurrentTypeReduction()

		if trw.isTypeDef() {
			field := trw.Fields[0]
			typeRed := ti.FieldTypeReduction(&typeReduction, 0)
			typeDependencies := field.t.ActualTypeDependencies(typeRed)

			for _, typeRw := range typeDependencies {
				if typeRw.cppLocalName != "" {
					hppInc.ns[typeRw] = CppIncludeInfo{componentId: typeRw.typeComponent, namespace: typeRw.tlName.Namespace}
				}
			}
			hpp.WriteString(fmt.Sprintf("using %s = %s;", trw.wr.cppLocalName, field.t.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, typeRed)))
		} else {
			hpp.WriteString("struct " + trw.wr.cppLocalName + " {\n")
			for i, field := range trw.Fields {
				hppIncByField := DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}

				typeRed := ti.FieldTypeReduction(&typeReduction, i)
				typeDeps := trw.Fields[i].t.ActualTypeDependencies(typeRed)
				for _, typeRw := range typeDeps {
					if typeRw.trw.IsWrappingType() {
						continue
					}
					hppIncByField.ns[typeRw] = CppIncludeInfo{componentId: typeRw.typeComponent, namespace: typeRw.tlName.Namespace}
				}

				fieldFullType := field.t.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, typeRed)
				fieldsMaskComment := ""
				//if field.fieldMask != nil {
				//	fieldsMaskComment = fmt.Sprintf(" // Conditional: %s.%d", formatNatArgCPP(trw.Fields, *field.fieldMask), field.BitNumber)
				//}
				if field.recursive {
					// TODO make better
					for includeType, includeInfo := range hppIncByField.ns {
						if includeInfo.componentId == trw.wr.typeComponent || includeType.cppLocalName == "" {
							delete(hppIncByField.ns, includeType)
						}
					}
					anyRecursive = true // requires destructor in cpp file
					hpp.WriteString(fmt.Sprintf("\tstd::shared_ptr<%s> %s{};%s\n", fieldFullType, field.cppName, fieldsMaskComment))
				} else {
					hpp.WriteString(fmt.Sprintf("\t%s %s%s;%s\n", fieldFullType, field.cppName, field.t.CPPDefaultInitializer(field.halfResolved, true), fieldsMaskComment))
				}
				for includeType, includeInfo := range hppIncByField.ns {
					hppInc.ns[includeType] = includeInfo
				}
				//hpp.WriteString(fmt.Sprintf("\t// DebugString: %s\n", field.resolvedType.DebugString()))
			}
			if anyRecursive { // && len(trw.cppArgs) != 0
				hpp.WriteString(fmt.Sprintf("\n\t~%s() {}\n", trw.wr.cppLocalName)) // TODO - move destructor to cpp
				// cppDet.WriteString(fmt.Sprintf("%s%s::~%s() {}\n", trw.wr.cppNamespaceQualifier, goLocalName, goLocalName))
			}
			if trw.wr.tlTag != 0 { // anonymous square brackets citizens or other exotic type
				hpp.WriteString(fmt.Sprintf(`
	// tl type info
	static constexpr uint32_t TL_TAG = 0x%08[2]x;
	static constexpr std::string_view TL_NAME = "%[1]s";

	uint32_t tl_tag() const { return 0x%08[2]x; }
	std::string_view tl_name() const { return "%[1]s"; }
`, trw.wr.tlName, trw.wr.tlTag))
			}
			if len(myArgsDecl) == 0 {
				// cppStartNamespace(cppDet, trw.wr.gen.RootCPPNamespaceElements)
				hpp.WriteString(fmt.Sprintf(`
	// basic serialization methods 
	bool write_json(std::ostream& s%[1]s) const;

	bool read(::%[6]s::tl_istream & s%[1]s) noexcept;
	bool write(::%[6]s::tl_ostream & s%[1]s) const noexcept;

	void read(::%[6]s::tl_throwable_istream & s%[1]s);
	void write(::%[6]s::tl_throwable_ostream & s%[1]s) const;
`,
					formatNatArgsDeclCPP(trw.wr.NatParams),
					trw.CPPTypeResettingCode(bytesVersion, "*this"),
					trw.CPPTypeReadingCode(bytesVersion, "*this", true, formatNatArgsAddNat(trw.wr.NatParams), true),
					trw.CPPTypeWritingCode(bytesVersion, "*this", true, formatNatArgsAddNat(trw.wr.NatParams), true),
					trw.wr.cppLocalName,
					trw.wr.gen.cppBasictlNamespace()))
				if trw.wr.tlTag != 0 { // anonymous square brackets citizens or other exotic type
					hpp.WriteString(fmt.Sprintf(`
	bool read_boxed(::%[5]s::tl_istream & s%[1]s) noexcept;
	bool write_boxed(::%[5]s::tl_ostream & s%[1]s) const noexcept;
	
	void read_boxed(::%[5]s::tl_throwable_istream & s%[1]s);
	void write_boxed(::%[5]s::tl_throwable_ostream & s%[1]s) const;
`,
						formatNatArgsDeclCPP(trw.wr.NatParams),
						trw.CPPTypeResettingCode(bytesVersion, "*this"),
						trw.CPPTypeReadingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true),
						trw.CPPTypeWritingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true),
						trw.wr.gen.cppBasictlNamespace()))
				}
				if trw.ResultType != nil {
					hppIncByResult := DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}

					typeRed := ti.ResultTypeReduction(&typeReduction)
					typeDeps := trw.ResultType.ActualTypeDependencies(typeRed)
					for _, typeRw := range typeDeps {
						if typeRw.trw.IsWrappingType() {
							continue
						}
						hppIncByResult.ns[typeRw] = CppIncludeInfo{componentId: typeRw.typeComponent, namespace: typeRw.tlName.Namespace}
					}
					for includeType, includeInfo := range hppIncByResult.ns {
						hppInc.ns[includeType] = includeInfo
					}
					hpp.WriteString(fmt.Sprintf(`
	// function methods and properties
	using ResultType = %[1]s;

	bool read_result(::%[2]s::tl_istream & s, %[1]s & result) const noexcept;
	bool write_result(::%[2]s::tl_ostream & s, const %[1]s & result) const noexcept;

	void read_result(::%[2]s::tl_throwable_istream & s, %[1]s & result) const;
	void write_result(::%[2]s::tl_throwable_ostream & s, const %[1]s & result) const;
`,
						trw.ResultType.CPPTypeStringInNamespaceHalfResolved2(bytesVersion, typeRed),
						trw.wr.gen.cppBasictlNamespace(),
					))
				}
				if len(trw.wr.NatParams) == 0 {
					hpp.WriteString(fmt.Sprintf(`
	friend std::ostream& operator<<(std::ostream& s, const %[1]s& rhs) {
		rhs.write_json(s);
		return s;
	}
`,
						trw.wr.cppLocalName))
				}
			}
			hpp.WriteString("};\n")
		}
		cppFinishNamespace(hpp, typeNamespace)
	}

	hppTmpInclude := DirectIncludesCPP{ns: map[*TypeRWWrapper]CppIncludeInfo{}}
	myFullType := trw.cppTypeStringInNamespace(bytesVersion, &hppTmpInclude)
	myFullTypeNoPrefix := strings.TrimPrefix(myFullType, "::") // Stupid C++ has sometimes problems with name resolution of definitions
	if hppDet != nil {
		cppStartNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)

		hppDet.WriteString(fmt.Sprintf(`
void %[1]sReset(%[2]s& item) noexcept;

bool %[1]sWriteJSON(std::ostream& s, const %[2]s& item%[3]s) noexcept;
bool %[1]sRead(::%[4]s::tl_istream & s, %[2]s& item%[3]s) noexcept; 
bool %[1]sWrite(::%[4]s::tl_ostream & s, const %[2]s& item%[3]s) noexcept;
`,
			goGlobalName,
			myFullType,
			formatNatArgsDeclCPP(trw.wr.NatParams),
			trw.wr.gen.cppBasictlNamespace()))

		if trw.wr.tlTag != 0 { // anonymous square brackets citizens or other exotic type
			hppDet.WriteString(fmt.Sprintf(`bool %[1]sReadBoxed(::%[4]s::tl_istream & s, %[2]s& item%[3]s);
bool %[1]sWriteBoxed(::%[4]s::tl_ostream & s, const %[2]s& item%[3]s);
`,
				goGlobalName,
				myFullType,
				formatNatArgsDeclCPP(trw.wr.NatParams),
				trw.wr.gen.cppBasictlNamespace()))
		}

		if trw.ResultType != nil {
			resultType := trw.ResultType.trw.cppTypeStringInNamespace(bytesVersion, &hppTmpInclude)
			hppDet.WriteString(fmt.Sprintf(`
bool %[1]sReadResult(::%[4]s::tl_istream & s, const %[2]s& item, %[3]s& result);
bool %[1]sWriteResult(::%[4]s::tl_ostream & s, const %[2]s& item, const %[3]s& result);
		`, goGlobalName, myFullType, resultType, trw.wr.gen.cppBasictlNamespace()))
		}

		cppFinishNamespace(hppDet, trw.wr.gen.DetailsCPPNamespaceElements)
		utils.AppendMap(hppTmpInclude.ns, &hppDetInc.ns)
	}

	if cppDet != nil {
		if !trw.isTypeDef() {
			if len(myArgsDecl) == 0 {
				cppDet.WriteString(fmt.Sprintf(`
bool %[5]s::write_json(std::ostream& s%[1]s)const {
%[6]s
	return true;
}

bool %[5]s::read(::%[8]s::tl_istream & s%[1]s) noexcept {
%[3]s
	s.sync();
	return true;
}

bool %[5]s::write(::%[8]s::tl_ostream & s%[1]s) const noexcept {
%[4]s
	s.sync();
	return true;
}

void %[5]s::read(::%[8]s::tl_throwable_istream & s%[1]s) {
	::%[8]s::tl_istream s2(s);
	this->read(s2%[7]s);
	s2.pass_data(s);
}

void %[5]s::write(::%[8]s::tl_throwable_ostream & s%[1]s) const {
	::%[8]s::tl_ostream s2(s);
	this->write(s2%[7]s);
	s2.pass_data(s);
}
`,
					formatNatArgsDeclCPP(trw.wr.NatParams),
					trw.CPPTypeResettingCode(bytesVersion, "*this"),
					trw.CPPTypeReadingCode(bytesVersion, "*this", true, formatNatArgsAddNat(trw.wr.NatParams), true),
					trw.CPPTypeWritingCode(bytesVersion, "*this", true, formatNatArgsAddNat(trw.wr.NatParams), true),
					myFullTypeNoPrefix,
					trw.CPPTypeWritingJsonCode(bytesVersion, "*this", true, formatNatArgsAddNat(trw.wr.NatParams), true),
					formatNatArgsCallCPP(trw.wr.NatParams),
					trw.wr.gen.cppBasictlNamespace(),
				))
				if trw.wr.tlTag != 0 {
					cppDet.WriteString(fmt.Sprintf(`
bool %[5]s::read_boxed(::%[7]s::tl_istream & s%[1]s) noexcept {
%[3]s
	s.sync();
	return true;
}

bool %[5]s::write_boxed(::%[7]s::tl_ostream & s%[1]s) const noexcept {
%[4]s
	s.sync();
	return true;
}

void %[5]s::read_boxed(::%[7]s::tl_throwable_istream & s%[1]s) {
	::%[7]s::tl_istream s2(s);
	this->read_boxed(s2%[6]s);
	s2.pass_data(s);
}

void %[5]s::write_boxed(::%[7]s::tl_throwable_ostream & s%[1]s) const {
	::%[7]s::tl_ostream s2(s);
	this->write_boxed(s2%[6]s);
	s2.pass_data(s);
}
`,
						formatNatArgsDeclCPP(trw.wr.NatParams),
						trw.CPPTypeResettingCode(bytesVersion, "*this"),
						trw.CPPTypeReadingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true),
						trw.CPPTypeWritingCode(bytesVersion, "*this", false, formatNatArgsAddNat(trw.wr.NatParams), true),
						myFullTypeNoPrefix,
						formatNatArgsCallCPP(trw.wr.NatParams),
						trw.wr.gen.cppBasictlNamespace()))
				}
			}
		}
		cppDet.WriteString(fmt.Sprintf(`
void %[7]s::%[1]sReset(%[2]s& item) noexcept {
	(void)item;
%[4]s}

bool %[7]s::%[1]sWriteJSON(std::ostream& s, const %[2]s& item%[3]s) noexcept {
	(void)s;
	(void)item;
%[8]s	return true;
}

bool %[7]s::%[1]sRead(::%[9]s::tl_istream & s, %[2]s& item%[3]s) noexcept {
	(void)s;
	(void)item;
%[5]s	return true;
}

bool %[7]s::%[1]sWrite(::%[9]s::tl_ostream & s, const %[2]s& item%[3]s) noexcept {
	(void)s;
	(void)item;
%[6]s	return true;
}
`,
			goGlobalName,
			myFullType,
			formatNatArgsDeclCPP(trw.wr.NatParams),
			trw.CPPResetFields(bytesVersion),
			trw.CPPReadFields(bytesVersion, hppDetInc, cppDetInc),
			trw.CPPWriteFields(bytesVersion),
			trw.wr.gen.DetailsCPPNamespace,
			trw.CPPWriteJsonFields(bytesVersion),
			trw.wr.gen.cppBasictlNamespace(),
		))

		if trw.wr.tlTag != 0 { // anonymous square brackets citizens or other exotic type
			cppDet.WriteString(fmt.Sprintf(`
bool %[7]s::%[1]sReadBoxed(::%[10]s::tl_istream & s, %[2]s& item%[3]s) {
	if (!s.nat_read_exact_tag(0x%08[9]x)) { return false; }
	return %[7]s::%[1]sRead(s, item%[8]s);
}

bool %[7]s::%[1]sWriteBoxed(::%[10]s::tl_ostream & s, const %[2]s& item%[3]s) {
	if (!s.nat_write(0x%08[9]x)) { return false; }
	return %[7]s::%[1]sWrite(s, item%[8]s);
}
`,
				goGlobalName,
				myFullType,
				formatNatArgsDeclCPP(trw.wr.NatParams),
				trw.CPPResetFields(bytesVersion),
				trw.CPPReadFields(bytesVersion, hppDetInc, cppDetInc),
				trw.CPPWriteFields(bytesVersion),
				trw.wr.gen.DetailsCPPNamespace,
				formatNatArgsCallCPP(trw.wr.NatParams),
				trw.wr.tlTag,
				trw.wr.gen.cppBasictlNamespace(),
			))
		}

		if trw.ResultType != nil {
			actualDep := trw.ResultType
			for {
				if strct, isStrct := actualDep.trw.(*TypeRWStruct); isStrct && strct.isUnwrapType() {
					actualDep = strct.Fields[0].t
				} else {
					break
				}
			}
			cppDetInc.ns[actualDep] = CppIncludeInfo{componentId: actualDep.typeComponent, namespace: actualDep.tlName.Namespace}
			resultType := trw.ResultType.trw.cppTypeStringInNamespace(bytesVersion, &hppTmpInclude)

			cppDet.WriteString(fmt.Sprintf(`
bool %[8]s::%[6]sReadResult(::%[9]s::tl_istream & s, const %[2]s& item, %[1]s& result) {
	(void)s;
	(void)item;
	(void)result;
%[3]s
	return true;
}
bool %[8]s::%[6]sWriteResult(::%[9]s::tl_ostream & s, const %[2]s& item, const %[1]s& result) {
	(void)s;
	(void)item;
	(void)result;
%[7]s
	return true;
}

bool %[2]s::read_result(::%[9]s::tl_istream & s, %[1]s & result) const noexcept {
	bool success = %[8]s::%[6]sReadResult(s, *this, result);
	s.sync();
	return success;
}
bool %[2]s::write_result(::%[9]s::tl_ostream & s, const %[1]s & result) const noexcept {
	bool success = %[8]s::%[6]sWriteResult(s, *this, result);
	s.sync();
	return success;
}

void %[2]s::read_result(::%[9]s::tl_throwable_istream & s, %[1]s & result) const {
	::%[9]s::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void %[2]s::write_result(::%[9]s::tl_throwable_ostream & s, const %[1]s & result) const {
	::%[9]s::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}
`,
				resultType,
				myFullTypeNoPrefix,
				trw.ResultType.trw.CPPTypeReadingCode(bytesVersion, "result", false, formatNatArgsCPP(trw.Fields, trw.ResultNatArgs), false),
				trw.ResultType.goGlobalName,
				joinWithCommas(formatNatArgsCPP(trw.Fields, trw.ResultNatArgs)),
				goGlobalName,
				trw.ResultType.trw.CPPTypeWritingCode(bytesVersion, "result", false, formatNatArgsCPP(trw.Fields, trw.ResultNatArgs), false),
				trw.wr.gen.DetailsCPPNamespace,
				trw.wr.gen.cppBasictlNamespace(),
			))
		}
	}
}

func (trw *TypeRWStruct) CPPResetFields(bytesVersion bool) string {
	var s strings.Builder
	if trw.isTypeDef() {
		field := trw.Fields[0]
		s.WriteString(
			field.t.trw.CPPTypeResettingCode(bytesVersion, "item") + "\n")
		return s.String()
	}
	for _, field := range trw.Fields {
		if field.recursive {
			s.WriteString(fmt.Sprintf(`
	if (item.%[1]s) {
		%[2]s
	}
`, field.cppName, field.t.trw.CPPTypeResettingCode(bytesVersion, fmt.Sprintf("(*item.%s)", field.cppName))))
		} else {
			s.WriteString(field.t.trw.CPPTypeResettingCode(bytesVersion, fmt.Sprintf("item.%s", field.cppName)) + "\n")
		}
	}
	return s.String()
}

func (trw *TypeRWStruct) CPPWriteFields(bytesVersion bool) string {
	var s strings.Builder
	if trw.isTypeDef() {
		field := trw.Fields[0]
		s.WriteString(
			field.t.trw.CPPTypeWritingCode(bytesVersion, "item",
				field.Bare(), formatNatArgsCPP(trw.Fields, field.natArgs), false) + "\n")
		return s.String()
	}
	for _, field := range trw.Fields {
		indent := 0
		if field.fieldMask != nil {
			s.WriteString(fmt.Sprintf("\tif ((%s & (1<<%d)) != 0) {\n\t", formatNatArgCPP(trw.Fields, *field.fieldMask), field.BitNumber))
			indent++
		}
		s.WriteString(strings.Repeat("\t", indent))
		s.WriteString(
			field.t.trw.CPPTypeWritingCode(bytesVersion, addAsterisk(field.recursive, fmt.Sprintf("item.%s", field.cppName)),
				field.Bare(), formatNatArgsCPP(trw.Fields, field.natArgs), false) + "\n")
		if field.fieldMask != nil {
			s.WriteString("\t}\n")
		}
	}
	return s.String()
}

func countFields(fields []Field) (count int) {
	for _, field := range fields {
		if field.t.IsTrueType() {
			if field.fieldMask == nil || !(field.fieldMask.isField || field.fieldMask.isArith) {
				continue
			}
		}
		count++
	}
	return
}

func (trw *TypeRWStruct) CPPWriteJsonFields(bytesVersion bool) string {
	var s strings.Builder
	if trw.isTypeDef() {
		field := trw.Fields[0]
		s.WriteString(
			field.t.trw.CPPTypeWritingJsonCode(bytesVersion, "item",
				field.Bare(), formatNatArgsCPP(trw.Fields, field.natArgs), false) + "\n")
		return s.String()
	}
	if trw.wr.IsTrueType() {
		return "	s << \"true\";\n"
	}

	fieldsCount := countFields(trw.Fields)
	if fieldsCount > 1 {
		s.WriteString(`	auto add_comma = false;
`)
	}
	s.WriteString(`	s << "{";
`)
	for i, field := range trw.Fields {
		if field.t.IsTrueType() {
			if field.fieldMask == nil || !(field.fieldMask.isField || field.fieldMask.isArith) {
				continue
			}
		}
		indent := 0
		if field.fieldMask != nil {
			s.WriteString(fmt.Sprintf("\tif ((%s & (1<<%d)) != 0) {\n", formatNatArgCPP(trw.Fields, *field.fieldMask), field.BitNumber))
			indent++
		}
		emptyCheck := field.t.trw.CPPTypeJSONEmptyCondition(bytesVersion, fmt.Sprintf("item.%s", field.cppName), field.recursive, formatNatArgsCPP(trw.Fields, field.natArgs))
		if field.fieldMask == nil && emptyCheck != "" {
			s.WriteString(fmt.Sprintf("%sif (%s) {\n", strings.Repeat("\t", indent+1), emptyCheck))
			indent++
		}
		if i != 0 {
			// append
			s.WriteString(fmt.Sprintf(`%[1]sif (add_comma) {
	%[1]ss << ",";
%[1]s}
`,
				strings.Repeat("\t", indent+1),
			))
		}
		if fieldsCount > 1 {
			s.WriteString(strings.Repeat("\t", indent+1))
			s.WriteString("add_comma = true;\n")
		}
		s.WriteString(fmt.Sprintf(`%ss << "\"%s\":";
`,
			strings.Repeat("\t", indent+1),
			field.originalName,
		))
		s.WriteString(strings.Repeat("\t", indent))
		s.WriteString(
			field.t.trw.CPPTypeWritingJsonCode(bytesVersion, addAsterisk(field.recursive, fmt.Sprintf("item.%s", field.cppName)),
				field.Bare(), formatNatArgsCPP(trw.Fields, field.natArgs), false) + "\n")
		if field.fieldMask == nil && emptyCheck != "" {
			s.WriteString(fmt.Sprintf("%s}\n", strings.Repeat("\t", indent)))
		}
		if field.fieldMask != nil {
			s.WriteString("\t}\n")
		}
	}
	s.WriteString(`	s << "}";
`)
	return s.String()
}

func (trw *TypeRWStruct) CPPReadFields(bytesVersion bool, hppDetInc *DirectIncludesCPP, cppDetInc *DirectIncludesCPP) string {
	var s strings.Builder
	if trw.isTypeDef() {
		field := trw.Fields[0]
		_ = field.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc) // only fill includes
		s.WriteString(
			field.t.trw.CPPTypeReadingCode(bytesVersion, "item",
				field.Bare(), formatNatArgsCPP(trw.Fields, field.natArgs),
				false) + "\n")

		return s.String()
	}
	for _, field := range trw.Fields {
		indent := 0
		if field.fieldMask != nil {
			s.WriteString(fmt.Sprintf("\tif ((%s & (1<<%d)) != 0) {\n", formatNatArgCPP(trw.Fields, *field.fieldMask), field.BitNumber))
			indent++
		}
		if field.recursive {
			s.WriteString(strings.Repeat("\t", indent))
			s.WriteString(fmt.Sprintf("\t"+`if (!item.%[2]s) { item.%[2]s = std::make_shared<%[1]s>(); }
`, field.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc), field.cppName))
		}
		if cppDetInc != nil {
			_ = field.t.CPPTypeStringInNamespace(bytesVersion, cppDetInc) // only fill includes
		}
		if hppDetInc != nil {
			_ = field.t.CPPTypeStringInNamespace(bytesVersion, hppDetInc) // only fill includes
		}
		s.WriteString(strings.Repeat("\t", indent))
		s.WriteString(
			field.t.trw.CPPTypeReadingCode(bytesVersion, addAsterisk(field.recursive, "item."+field.cppName),
				field.Bare(), formatNatArgsCPP(trw.Fields, field.natArgs),
				false) + "\n")
		if field.fieldMask != nil {
			// TODO - in case of recursive field, check for nil ptr
			s.WriteString("\t} else {\n")
			if field.recursive {
				s.WriteString(fmt.Sprintf("\t\tif (item.%s) {\n", field.cppName))
				s.WriteString(fmt.Sprintf("\t\t%s\n", field.t.trw.CPPTypeResettingCode(bytesVersion, addAsterisk(field.recursive, fmt.Sprintf("item.%s", field.cppName)))))
				s.WriteString("\t\t}\n")
			} else {
				s.WriteString(fmt.Sprintf("\t\t%s\n", field.t.trw.CPPTypeResettingCode(bytesVersion, addAsterisk(field.recursive, fmt.Sprintf("item.%s", field.cppName)))))
			}
			s.WriteString("\t}\n")
		}
	}
	return s.String()
}
