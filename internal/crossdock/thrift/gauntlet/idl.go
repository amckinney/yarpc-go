// Code generated by thriftrw v1.2.0
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gauntlet

import "go.uber.org/thriftrw/thriftreflect"

var ThriftModule = &thriftreflect.ThriftModule{Name: "gauntlet", Package: "go.uber.org/yarpc/internal/crossdock/thrift/gauntlet", FilePath: "gauntlet.thrift", SHA1: "f308242071d9c9c31fb8faa1e1a79eeefa46316d", Raw: rawIDL}

const rawIDL = "/**\n * Thrift file used by Apache Thrift's test suite, DO NOT EDIT!\n * Lifted from: https://thrift.apache.org/docs/idl\n */\n\nnamespace c_glib TTest\nnamespace java thrift.test\nnamespace cpp thrift.test\nnamespace rb Thrift.Test\nnamespace perl ThriftTest\nnamespace csharp Thrift.Test\nnamespace js ThriftTest\nnamespace st ThriftTest\nnamespace py ThriftTest\nnamespace py.twisted ThriftTest\nnamespace go thrifttest\nnamespace php ThriftTest\nnamespace delphi Thrift.Test\nnamespace cocoa ThriftTest\nnamespace lua ThriftTest\n\n// Presence of namespaces and sub-namespaces for which there is\n// no generator should compile with warnings only\nnamespace noexist ThriftTest\nnamespace cpp.noexist ThriftTest\n\nnamespace * thrift.test\n\n/**\n * Docstring!\n */\nenum Numberz\n{\n  ONE = 1,\n  TWO,\n  THREE,\n  FIVE = 5,\n  SIX,\n  EIGHT = 8\n}\n\nconst Numberz myNumberz = Numberz.ONE;\n// the following is expected to fail:\n// const Numberz urNumberz = ONE;\n\ntypedef i64 UserId\n\nstruct Bonk\n{\n  1: optional string message,\n  2: optional i32 type\n}\n\ntypedef map<string,Bonk> MapType\n\nstruct Bools {\n  1: optional bool im_true,\n  2: optional bool im_false,\n}\n\nstruct Xtruct\n{\n  1:  optional string string_thing,\n  4:  optional byte   byte_thing,\n  9:  optional i32    i32_thing,\n  11: optional i64    i64_thing\n}\n\nstruct Xtruct2\n{\n  1: optional byte   byte_thing,\n  2: optional Xtruct struct_thing,\n  3: optional i32    i32_thing\n}\n\nstruct Xtruct3\n{\n  1:  optional string string_thing,\n  4:  optional i32    changed,\n  9:  optional i32    i32_thing,\n  11: optional i64    i64_thing\n}\n\n\nstruct Insanity\n{\n  1: optional map<Numberz, UserId> userMap,\n  2: optional list<Xtruct> xtructs\n}\n\nstruct CrazyNesting {\n  1: optional string string_field,\n  2: optional set<Insanity> set_field,\n  3: required list< map<set<i32>,map<i32,set<list<map<Insanity,string>>>>>> list_field,\n  4: optional binary binary_field\n}\n\nexception Xception {\n  1: optional i32 errorCode,\n  2: optional string message\n}\n\nexception Xception2 {\n  1: optional i32 errorCode,\n  2: optional Xtruct struct_thing\n}\n\nstruct EmptyStruct {}\n\nstruct OneField {\n  1: optional EmptyStruct field\n}\n\nservice ThriftTest\n{\n  /**\n   * Prints \"testVoid()\" and returns nothing.\n   */\n  void         testVoid(),\n\n  /**\n   * Prints 'testString(\"%s\")' with thing as '%s'\n   * @param string thing - the string to print\n   * @return string - returns the string 'thing'\n   */\n  string       testString(1: string thing),\n\n  /**\n   * Prints 'testByte(\"%d\")' with thing as '%d'\n   * @param byte thing - the byte to print\n   * @return byte - returns the byte 'thing'\n   */\n  byte         testByte(1: byte thing),\n\n  /**\n   * Prints 'testI32(\"%d\")' with thing as '%d'\n   * @param i32 thing - the i32 to print\n   * @return i32 - returns the i32 'thing'\n   */\n  i32          testI32(1: i32 thing),\n\n  /**\n   * Prints 'testI64(\"%d\")' with thing as '%d'\n   * @param i64 thing - the i64 to print\n   * @return i64 - returns the i64 'thing'\n   */\n  i64          testI64(1: i64 thing),\n\n  /**\n   * Prints 'testDouble(\"%f\")' with thing as '%f'\n   * @param double thing - the double to print\n   * @return double - returns the double 'thing'\n   */\n  double       testDouble(1: double thing),\n\n  /**\n   * Prints 'testBinary(\"%s\")' where '%s' is a hex-formatted string of thing's data\n   * @param binary  thing - the binary data to print\n   * @return binary  - returns the binary 'thing'\n   */\n  binary       testBinary(1: binary thing),\n\n  /**\n   * Prints 'testStruct(\"{%s}\")' where thing has been formatted into a string of comma separated values\n   * @param Xtruct thing - the Xtruct to print\n   * @return Xtruct - returns the Xtruct 'thing'\n   */\n  Xtruct       testStruct(1: Xtruct thing),\n\n  /**\n   * Prints 'testNest(\"{%s}\")' where thing has been formatted into a string of the nested struct\n   * @param Xtruct2 thing - the Xtruct2 to print\n   * @return Xtruct2 - returns the Xtruct2 'thing'\n   */\n  Xtruct2      testNest(1: Xtruct2 thing),\n\n  /**\n   * Prints 'testMap(\"{%s\")' where thing has been formatted into a string of  'key => value' pairs\n   *  separated by commas and new lines\n   * @param map<i32,i32> thing - the map<i32,i32> to print\n   * @return map<i32,i32> - returns the map<i32,i32> 'thing'\n   */\n  map<i32,i32> testMap(1: map<i32,i32> thing),\n\n  /**\n   * Prints 'testStringMap(\"{%s}\")' where thing has been formatted into a string of  'key => value' pairs\n   *  separated by commas and new lines\n   * @param map<string,string> thing - the map<string,string> to print\n   * @return map<string,string> - returns the map<string,string> 'thing'\n   */\n  map<string,string> testStringMap(1: map<string,string> thing),\n\n  /**\n   * Prints 'testSet(\"{%s}\")' where thing has been formatted into a string of  values\n   *  separated by commas and new lines\n   * @param set<i32> thing - the set<i32> to print\n   * @return set<i32> - returns the set<i32> 'thing'\n   */\n  set<i32>     testSet(1: set<i32> thing),\n\n  /**\n   * Prints 'testList(\"{%s}\")' where thing has been formatted into a string of  values\n   *  separated by commas and new lines\n   * @param list<i32> thing - the list<i32> to print\n   * @return list<i32> - returns the list<i32> 'thing'\n   */\n  list<i32>    testList(1: list<i32> thing),\n\n  /**\n   * Prints 'testEnum(\"%d\")' where thing has been formatted into it's numeric value\n   * @param Numberz thing - the Numberz to print\n   * @return Numberz - returns the Numberz 'thing'\n   */\n  Numberz      testEnum(1: Numberz thing),\n\n  /**\n   * Prints 'testTypedef(\"%d\")' with thing as '%d'\n   * @param UserId thing - the UserId to print\n   * @return UserId - returns the UserId 'thing'\n   */\n  UserId       testTypedef(1: UserId thing),\n\n  /**\n   * Prints 'testMapMap(\"%d\")' with hello as '%d'\n   * @param i32 hello - the i32 to print\n   * @return map<i32,map<i32,i32>> - returns a dictionary with these values:\n   *   {-4 => {-4 => -4, -3 => -3, -2 => -2, -1 => -1, }, 4 => {1 => 1, 2 => 2, 3 => 3, 4 => 4, }, }\n   */\n  map<i32,map<i32,i32>> testMapMap(1: i32 hello),\n\n  /**\n   * So you think you've got this all worked, out eh?\n   *\n   * Creates a the returned map with these values and prints it out:\n   *   { 1 => { 2 => argument,\n   *            3 => argument,\n   *          },\n   *     2 => { 6 => <empty Insanity struct>, },\n   *   }\n   * @return map<UserId, map<Numberz,Insanity>> - a map with the above values\n   */\n  map<UserId, map<Numberz,Insanity>> testInsanity(1: Insanity argument),\n\n  /**\n   * Prints 'testMulti()'\n   * @param byte arg0 -\n   * @param i32 arg1 -\n   * @param i64 arg2 -\n   * @param map<i16, string> arg3 -\n   * @param Numberz arg4 -\n   * @param UserId arg5 -\n   * @return Xtruct - returns an Xtruct with string_thing = \"Hello2, byte_thing = arg0, i32_thing = arg1\n   *    and i64_thing = arg2\n   */\n  Xtruct testMulti(1: byte arg0, 2: i32 arg1, 3: i64 arg2, 4: map<i16, string> arg3, 5: Numberz arg4, 6: UserId arg5),\n\n  /**\n   * Print 'testException(%s)' with arg as '%s'\n   * @param string arg - a string indication what type of exception to throw\n   * if arg == \"Xception\" throw Xception with errorCode = 1001 and message = arg\n   * elsen if arg == \"TException\" throw TException\n   * else do not throw anything\n   */\n  void testException(1: string arg) throws(1: Xception err1),\n\n  /**\n   * Print 'testMultiException(%s, %s)' with arg0 as '%s' and arg1 as '%s'\n   * @param string arg - a string indication what type of exception to throw\n   * if arg0 == \"Xception\" throw Xception with errorCode = 1001 and message = \"This is an Xception\"\n   * elsen if arg0 == \"Xception2\" throw Xception2 with errorCode = 2002 and message = \"This is an Xception2\"\n   * else do not throw anything\n   * @return Xtruct - an Xtruct with string_thing = arg1\n   */\n  Xtruct testMultiException(1: string arg0, 2: string arg1) throws(1: Xception err1, 2: Xception2 err2)\n\n  /**\n   * Print 'testOneway(%d): Sleeping...' with secondsToSleep as '%d'\n   * sleep 'secondsToSleep'\n   * Print 'testOneway(%d): done sleeping!' with secondsToSleep as '%d'\n   * @param i32 secondsToSleep - the number of seconds to sleep\n   */\n  oneway void testOneway(1:i32 secondsToSleep)\n}\n\nservice SecondService\n{\n  void blahBlah()\n  /**\n   * Prints 'testString(\"%s\")' with thing as '%s'\n   * @param string thing - the string to print\n   * @return string - returns the string 'thing'\n   */\n  string       secondtestString(1: string thing),\n}\n\nstruct VersioningTestV1 {\n       1: optional i32 begin_in_both,\n       3: optional string old_string,\n       12: optional i32 end_in_both\n}\n\nstruct VersioningTestV2 {\n       1: optional i32 begin_in_both,\n\n       2: optional i32 newint,\n       3: optional byte newbyte,\n       4: optional i16 newshort,\n       5: optional i64 newlong,\n       6: optional double newdouble\n       7: optional Bonk newstruct,\n       8: optional list<i32> newlist,\n       9: optional set<i32> newset,\n       10: optional map<i32, i32> newmap,\n       11: optional string newstring,\n       12: optional i32 end_in_both\n}\n\nstruct ListTypeVersioningV1 {\n       1: optional list<i32> myints;\n       2: optional string hello;\n}\n\nstruct ListTypeVersioningV2 {\n       1: optional list<string> strings;\n       2: optional string hello;\n}\n\nstruct GuessProtocolStruct {\n  7: optional map<string,string> map_field,\n}\n\nstruct LargeDeltas {\n  1: optional Bools b1,\n  10: optional Bools b10,\n  100: optional Bools b100,\n  500: optional bool check_true,\n  1000: optional Bools b1000,\n  1500: optional bool check_false,\n  2000: optional VersioningTestV2 vertwo2000,\n  2500: optional set<string> a_set2500,\n  3000: optional VersioningTestV2 vertwo3000,\n  4000: optional list<i32> big_numbers\n}\n\nstruct NestedListsI32x2 {\n  1: optional list<list<i32>> integerlist\n}\nstruct NestedListsI32x3 {\n  1: optional list<list<list<i32>>> integerlist\n}\nstruct NestedMixedx2 {\n  1: optional list<set<i32>> int_set_list\n  2: optional map<i32,set<string>> map_int_strset\n  3: optional list<map<i32,set<string>>> map_int_strset_list\n}\nstruct ListBonks {\n  1: optional list<Bonk> bonk\n}\nstruct NestedListsBonk {\n  1: optional list<list<list<Bonk>>> bonk\n}\n\nstruct BoolTest {\n  1: optional bool b = true;\n  2: optional string s = \"true\";\n}\n\nstruct StructA {\n  1: required string s;\n}\n\nstruct StructB {\n  1: optional StructA aa;\n  2: required StructA ab;\n}\n"
