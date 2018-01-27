/*
Navicat PGSQL Data Transfer

Source Server         : LOCAL
Source Server Version : 90504
Source Host           : localhost:5432
Source Database       : fixman
Source Schema         : public

Target Server Type    : PGSQL
Target Server Version : 90504
File Encoding         : 65001

Date: 2018-01-27 09:30:02
*/


-- ----------------------------
-- Table structure for provinces
-- ----------------------------
DROP TABLE IF EXISTS "public"."provinces";
CREATE TABLE "public"."provinces" (
"i_d" int4 DEFAULT nextval('provinces_i_d_seq'::regclass) NOT NULL,
"name" text COLLATE "default" DEFAULT ''::text NOT NULL,
"created_at" timestamptz(6),
"updated_at" timestamptz(6)
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Records of provinces
-- ----------------------------
INSERT INTO "public"."provinces" VALUES ('1', 'กระบี่', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('2', 'กรุงเทพมหานคร', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('3', 'กาญจนบุรี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('4', 'กาฬสินธุ์', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('5', 'กำแพงเพชร', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('6', 'ขอนแก่น', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('7', 'จันทบุรี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('8', 'ฉะเชิงเทรา', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('9', 'ชลบุรี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('10', 'ชัยนาท', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('11', 'ชัยภูมิ', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('12', 'ชุมพร', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('13', 'เชียงราย', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('14', 'เชียงใหม่', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('15', 'ตรัง', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('16', 'ตราด', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('17', 'ตาก', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('18', 'นครนายก', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('19', 'นครปฐม', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('20', 'นครพนม', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('21', 'นครราชสีมา', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('22', 'นครศรีธรรมราช', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('23', 'นครสวรรค์', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('24', 'นนทบุรี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('25', 'นราธิวาส', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('26', 'น่าน', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('27', 'บุรีรัมย์', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('28', 'ปทุมธานี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('29', 'ประจวบคีรีขันธ์', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('30', 'ปราจีนบุรี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('31', 'ปัตตานี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('32', 'พระนครศรีอยุธยา', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('33', 'พะเยา', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('34', 'พังงา', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('35', 'พัทลุง', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('36', 'พิจิตร', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('37', 'พิษณุโลก', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('38', 'เพชรบุรี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('39', 'เพชรบูรณ์', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('40', 'แพร่', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('41', 'ภูเก็ต', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('42', 'มหาสารคาม', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('43', 'มุกดาหาร', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('44', 'แม่ฮ่องสอน', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('45', 'ยโสธร', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('46', 'ยะลา', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('47', 'ร้อยเอ็ด', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('48', 'ระนอง', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('49', 'ระยอง', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('50', 'ราชบุรี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('51', 'ลพบุรี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('52', 'เลย', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('53', 'ลำปาง', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('54', 'ลำพูน', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('55', 'ศีรสะเกษ', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('56', 'สกลนคร', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('57', 'สงขลา', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('58', 'สตูล', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('59', 'สมุทรปราการ', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('60', 'สมุทรสงคราม', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('61', 'สมุทรสาคร', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('62', 'สระแก้ว', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('63', 'สระบุรี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('64', 'สิงห์บุรี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('65', 'สุโขทัย', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('66', 'สุพรรณบุรี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('67', 'สุราษฎร์ธานี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('68', 'สุรินทร์', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('69', 'หนองคาย', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('70', 'หนองบัวลำภู', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('71', 'อ่างทอง', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('72', 'อำนาจเจริญ', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('73', 'อุดรธานี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('74', 'อุตรดิตถ์', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('75', 'อุทัยธานี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('76', 'อุบลราชธานี', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');
INSERT INTO "public"."provinces" VALUES ('77', 'บึงกาฬ', '2018-01-08 10:08:00+07', '2018-01-08 10:08:00+07');

-- ----------------------------
-- Alter Sequences Owned By 
-- ----------------------------

-- ----------------------------
-- Primary Key structure for table provinces
-- ----------------------------
ALTER TABLE "public"."provinces" ADD PRIMARY KEY ("i_d");
