/*
 Navicat Premium Data Transfer

 Source Server         : GaussDB
 Source Server Type    : GaussDB
 Source Server Version : 90204 (90204)
 Source Host           : 118.196.21.42:15432
 Source Catalog        : cardgame
 Source Schema         : public

 Target Server Type    : GaussDB
 Target Server Version : 90204 (90204)
 File Encoding         : 65001

 Date: 16/06/2025 12:48:34
*/


-- ----------------------------
-- Table structure for cards
-- ----------------------------
DROP TABLE IF EXISTS "public"."cards";
CREATE TABLE "public"."cards" (
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "type" varchar(255) COLLATE "pg_catalog"."default",
  "energy" int4,
  "target" varchar(255) COLLATE "pg_catalog"."default",
  "block" int4,
  "damage" int4,
  "power" text COLLATE "pg_catalog"."default",
  "action" text COLLATE "pg_catalog"."default",
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "image" varchar(255) COLLATE "pg_catalog"."default",
  "upgrade" int4,
  "id" int4
)
WITH (orientation=ROW)
;

-- ----------------------------
-- Records of cards
-- ----------------------------
INSERT INTO "public"."cards" VALUES ('仪式', 'skill', 2, 'player', 0, 0, NULL, '["removePlayerDebuffs"]', '移除所有负面状态', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('仪式(S)', 'skill', 2, 'player', 10, 0, NULL, '["removePlayerDebuffs"]', '移除所有负面状态，得到10点格挡值', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('吸血', 'attack', 3, 'enemy', 0, 0, NULL, '["addRegenEqualToAllDamage"]', '造成2点伤害回复2点生命值', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('吸血(S)', 'attack', 2, 'enemy', 0, 0, NULL, '["addRegenEqualToAllDamage"]', '造成2点伤害回复2点生命值', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('恐吓', 'skill', 0, 'enemy', 0, 0, '{"vulnerable": 1}', NULL, '造成1层易伤', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('恐吓(S)', 'skill', 0, 'enemy', 0, 0, '{"vulnerable": 2}', NULL, '造成2层易伤', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('有毒物质', 'attack', 1, 'enemy', 0, 0, '{"poison": 5}', NULL, '造成5层中毒', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('有毒物质(S)', 'attack', 1, 'enemy', 0, 0, '{"poison": 7}', NULL, '造成7点中毒', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('棍棒打击', 'attack', 3, 'enemy', 0, 24, NULL, NULL, '造成24点伤害', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('棍棒打击(S)', 'attack', 3, 'enemy', 0, 36, NULL, NULL, '造成36点伤害', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('灵魂药水', 'attack', 1, 'enemy', 0, 0, '{"weak": 2, "vulnerable": 2}', '["removeHealth"]', '造成2层虚弱和2层易伤，失去3点生命', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('灵魂药水(S)', 'attack', 1, 'enemy', 0, 0, '{"weak": 3, "vulnerable": 3}', '["removeHealth"]', '造成3层虚弱和3层易伤，失去2点生命', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('猛击', 'attack', 2, 'enemy', 0, 8, '{"vulnerable": 2}', NULL, '造成8点伤害，造成2层易伤', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('猛击(S)', 'attack', 2, 'enemy', 0, 10, '{"vulnerable": 3}', NULL, '造成10点伤害，造成3层易伤', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('砍倒', 'attack', 1, 'enemy', 0, 8, NULL, NULL, '造成8点伤害', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('砍倒(S)', 'attack', 1, 'enemy', 0, 11, NULL, NULL, '造成11点伤害', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('突袭', 'attack', 1, 'enemy', 0, 7, '{"weak": 1}', NULL, '造成7点伤害，造成1层虚弱', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('突袭(S)', 'attack', 1, 'enemy', 0, 8, '{"weak": 2}', NULL, '造成8点伤害，造成2层虚弱', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('紧急', 'skill', 0, 'player', 0, 0, NULL, '["addEnergyToPlayer"]', '得到1点能量', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('紧急(S)', 'skill', 0, 'player', 5, 0, NULL, '["addEnergyToPlayer"]', '得到1点能量,得到5点格挡值', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('罢工', 'attack', 1, 'enemy', 0, 9, NULL, '["drawCards"]', '造成9点伤害，抽1张牌', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('罢工(S)', 'attack', 1, 'enemy', 0, 13, NULL, '["drawCards"]', '造成13点伤害，抽1张牌', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('肾上腺素', 'skill', 0, 'player', 0, 0, NULL, '["drawCards", "addEnergyToPlayer"]', '得到1点能量，抽1张牌', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('肾上腺素(S)', 'skill', 0, 'player', 0, 2, NULL, '["drawCards", "addEnergyToPlayer"]', '造成2点伤害，得到1点能量，抽1张牌', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('袭击', 'attack', 1, 'enemy', 0, 6, NULL, NULL, '造成6点伤害', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('袭击(S)', 'attack', 1, 'enemy', 0, 9, NULL, NULL, '造成9点伤害', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('身体撞击', 'attack', 1, 'enemy', 0, 0, NULL, '["dealDamageEqualToBlock"]', '造成与格挡值相同的伤害', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('身体撞击(S)', 'attack', 0, 'enemy', 0, 0, NULL, '["dealDamageEqualToBlock"]', '造成与格挡值相同的伤害', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('钢铁撞击', 'attack', 1, 'enemy', 5, 5, NULL, NULL, '造成5点伤害，得到5点格挡值', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('钢铁撞击(S)', 'attack', 1, 'enemy', 7, 7, NULL, NULL, '造成7点伤害，得到7点格挡值', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('防御', 'skill', 1, 'player', 5, 0, NULL, NULL, '得到5点格挡值', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('防御(S)', 'skill', 1, 'player', 8, 0, NULL, NULL, '得到8点格挡值', NULL, 1, NULL);
INSERT INTO "public"."cards" VALUES ('雷霆一击', 'attack', 1, 'enemy', 0, 4, '{"vulnerable": 1}', NULL, '造成4点伤害，造成1层易伤', NULL, 0, NULL);
INSERT INTO "public"."cards" VALUES ('雷霆一击(S)', 'attack', 1, 'enemy', 0, 6, '{"vulnerable": 1}', NULL, '造成6点伤害，造成1层易伤', NULL, 1, NULL);

-- ----------------------------
-- Table structure for leaderboards
-- ----------------------------
DROP TABLE IF EXISTS "public"."leaderboards";
CREATE TABLE "public"."leaderboards" (
  "playerID" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "counts" int4,
  "id" int4
)
WITH (orientation=ROW)
;

-- ----------------------------
-- Records of leaderboards
-- ----------------------------
INSERT INTO "public"."leaderboards" VALUES ('admin', 1, NULL);
INSERT INTO "public"."leaderboards" VALUES ('test', 2, NULL);
INSERT INTO "public"."leaderboards" VALUES ('zk', 21, NULL);

-- ----------------------------
-- Table structure for monsters
-- ----------------------------
DROP TABLE IF EXISTS "public"."monsters";
CREATE TABLE "public"."monsters" (
  "HP" int4,
  "block" int4,
  "power" text COLLATE "pg_catalog"."default",
  "image" varchar(255) COLLATE "pg_catalog"."default",
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "id" int4,
  "Type" varchar(255) COLLATE "pg_catalog"."default",
  "actionName" text COLLATE "pg_catalog"."default",
  "actionValue" text COLLATE "pg_catalog"."default"
)
WITH (orientation=ROW)
;

-- ----------------------------
-- Records of monsters
-- ----------------------------
INSERT INTO "public"."monsters" VALUES (100, 0, NULL, NULL, 'boss1号', 10, 'boss', '["weak", "block", "damage"]', '[2, 6, 16]');
INSERT INTO "public"."monsters" VALUES (62, 0, NULL, NULL, 'boss2号', 11, 'boss', '["damage", "damage", "damage", "damage", "damage", "damage", "damage", "damage"]', '[5, 8, 12, 17, 23, 30, 38, 45]');
INSERT INTO "public"."monsters" VALUES (8, 0, NULL, NULL, '怪兽1号', 1, 'monster', '["block", "damage"]', '[9, 7]');
INSERT INTO "public"."monsters" VALUES (14, 0, NULL, NULL, '怪兽2号', 2, 'monster', '["block", "damage"]', '[10, 8]');
INSERT INTO "public"."monsters" VALUES (20, 0, NULL, NULL, '怪兽3号', 3, 'monster', '["block", "damage"]', '[10, 8]');
INSERT INTO "public"."monsters" VALUES (33, 0, NULL, NULL, '怪兽4号', 4, 'monster', '["weak", "damage", "vulnerable"]', '[1, 10, 1]');
INSERT INTO "public"."monsters" VALUES (40, 0, NULL, NULL, '怪兽5号', 5, 'monster', '["block", "damage"]', '[5, 11]');
INSERT INTO "public"."monsters" VALUES (50, 0, NULL, NULL, '怪兽6号', 6, 'monster', '["weak", "damage"]', '[1, 9]');
INSERT INTO "public"."monsters" VALUES (46, 0, NULL, NULL, '精英1号', 7, 'elites', '["block", "damage"]', '[6, 12]');
INSERT INTO "public"."monsters" VALUES (60, 0, NULL, NULL, '精英2号', 8, 'elites', '["weak", "damage", "vulnerable"]', '[1, 12, 1]');
INSERT INTO "public"."monsters" VALUES (70, 0, NULL, NULL, '精英3号', 9, 'elites', '["block", "damage"]', '[5, 16]');

-- ----------------------------
-- Table structure for userconfigs
-- ----------------------------
DROP TABLE IF EXISTS "public"."userconfigs";
CREATE TABLE "public"."userconfigs" (
  "playerID" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "cards" varchar(255) COLLATE "pg_catalog"."default",
  "ladder" varchar(1000) COLLATE "pg_catalog"."default",
  "playerHP" int4,
  "playerEnergy" int4,
  "image" varchar(255) COLLATE "pg_catalog"."default",
  "id" int4 DEFAULT 0
)
WITH (orientation=ROW)
;

-- ----------------------------
-- Records of userconfigs
-- ----------------------------
INSERT INTO "public"."userconfigs" VALUES ('adyu', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击","吸血"]', '[{"ID":0,"Name":"怪兽1号","Visit":1},{"ID":1,"Name":"怪兽3号","Visit":0},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英2号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss1号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('qweq', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽4号","Visit":0},{"ID":2,"Name":"怪兽5号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英2号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('qwer', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽4号","Visit":0},{"ID":2,"Name":"怪兽5号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英2号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('qwert', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽3号","Visit":0},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss1号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('zkk', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽2号","Visit":0},{"ID":1,"Name":"怪兽3号","Visit":0},{"ID":2,"Name":"怪兽5号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('zkz23zk', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽3号","Visit":0},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英2号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('zxc', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽3号","Visit":0},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英2号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss1号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('098765432', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽4号","Visit":0},{"ID":2,"Name":"怪兽5号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英2号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, NULL);
INSERT INTO "public"."userconfigs" VALUES ('10', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽2号","Visit":0},{"ID":2,"Name":"怪兽3号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss1号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('123', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽4号","Visit":0},{"ID":2,"Name":"怪兽5号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英2号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss1号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('12345', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽3号","Visit":0},{"ID":2,"Name":"怪兽5号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英2号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss1号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('a', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽2号","Visit":0},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英2号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss1号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('admin', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击","砍倒"]', '[{"ID":0,"Name":"怪兽1号","Visit":1},{"ID":1,"Name":"怪兽2号","Visit":0},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英2号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('iuih', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽3号","Visit":0},{"ID":1,"Name":"怪兽4号","Visit":0},{"ID":2,"Name":"怪兽5号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英2号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss1号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('player1', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击","钢铁撞击(S)","钢铁撞击","紧急(S)","仪式"]', '[{"ID":0,"Name":"怪兽2号","Visit":1},{"ID":1,"Name":"怪兽4号","Visit":1},{"ID":2,"Name":"怪兽5号","Visit":1},{"ID":3,"Name":"campfire","Visit":1},{"ID":4,"Name":"精英2号","Visit":1},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 70, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('popwiqpe', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽2号","Visit":0},{"ID":2,"Name":"怪兽5号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英2号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss1号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('zk', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击","棍棒打击(S)","紧急"]', '[{"ID":0,"Name":"怪兽1号","Visit":1},{"ID":1,"Name":"怪兽3号","Visit":1},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英2号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 58, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('zxcv', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽2号","Visit":0},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('zzkk', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽2号","Visit":0},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('zzzk', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击","有毒物质","吸血(S)"]', '[{"ID":0,"Name":"怪兽1号","Visit":1},{"ID":1,"Name":"怪兽2号","Visit":1},{"ID":2,"Name":"怪兽3号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英2号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('123456789', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击","钢铁撞击","灵魂药水(S)"]', '[{"ID":0,"Name":"怪兽1号","Visit":1},{"ID":1,"Name":"怪兽3号","Visit":1},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 69, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('0987654321', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽3号","Visit":0},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英2号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss1号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('09876543', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽2号","Visit":0},{"ID":2,"Name":"怪兽5号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('112233', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽2号","Visit":0},{"ID":1,"Name":"怪兽3号","Visit":0},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英2号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('1223', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽2号","Visit":0},{"ID":2,"Name":"怪兽5号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英2号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('1231', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽1号","Visit":0},{"ID":1,"Name":"怪兽2号","Visit":0},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英2号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss1号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('1890', '["仪式","吸血","恐吓","有毒物质","棍棒打击(S)","灵魂药水","猛击","砍倒","突袭","肾上腺素","身体撞击","防御","棍棒打击(S)","突袭(S)"]', '[{"ID":0,"Name":"怪兽1号","Visit":1},{"ID":1,"Name":"怪兽2号","Visit":1},{"ID":2,"Name":"怪兽4号","Visit":1},{"ID":3,"Name":"campfire","Visit":1},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英2号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('0987', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽2号","Visit":0},{"ID":1,"Name":"怪兽3号","Visit":0},{"ID":2,"Name":"怪兽4号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英3号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);
INSERT INTO "public"."userconfigs" VALUES ('test', '["袭击","袭击","袭击","袭击","袭击","防御","防御","防御","防御","猛击"]', '[{"ID":0,"Name":"怪兽3号","Visit":0},{"ID":1,"Name":"怪兽4号","Visit":0},{"ID":2,"Name":"怪兽5号","Visit":0},{"ID":3,"Name":"campfire","Visit":0},{"ID":4,"Name":"精英1号","Visit":0},{"ID":5,"Name":"精英2号","Visit":0},{"ID":6,"Name":"campfire","Visit":0},{"ID":7,"Name":"boss2号","Visit":0}]', 72, 3, NULL, 0);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "username" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "id" int4
)
WITH (orientation=ROW)
;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "public"."users" VALUES ('admin', 'admin', 1);
INSERT INTO "public"."users" VALUES ('player1', '123456', 2);
INSERT INTO "public"."users" VALUES ('zk', '123456', 3);
INSERT INTO "public"."users" VALUES ('zkz23zk', '123456', 4);
INSERT INTO "public"."users" VALUES ('zkzk', '123456', 5);
INSERT INTO "public"."users" VALUES ('zkzzk', '123456', 6);
INSERT INTO "public"."users" VALUES ('zzkk', '123456', 7);
INSERT INTO "public"."users" VALUES ('zkk', '123456', 8);
INSERT INTO "public"."users" VALUES ('qweq', '123456', 9);
INSERT INTO "public"."users" VALUES ('adyu', '123456', 10);
INSERT INTO "public"."users" VALUES ('qwer', '123456', 11);
INSERT INTO "public"."users" VALUES ('123', '123456', 12);
INSERT INTO "public"."users" VALUES ('12345', '123456', 13);
INSERT INTO "public"."users" VALUES ('zxcv', '123456', 14);
INSERT INTO "public"."users" VALUES ('10', '123456', 15);
INSERT INTO "public"."users" VALUES ('zxc', '123456', 16);
INSERT INTO "public"."users" VALUES ('qwert', '123456', 17);
INSERT INTO "public"."users" VALUES ('test', '123456', 18);
INSERT INTO "public"."users" VALUES ('a', '123456', 19);
INSERT INTO "public"."users" VALUES ('popwiqpe', '123456', 20);
INSERT INTO "public"."users" VALUES ('iuih', '123456', 21);
INSERT INTO "public"."users" VALUES ('zzzk', '123456', 22);
INSERT INTO "public"."users" VALUES ('123456789', '123456789', 23);
INSERT INTO "public"."users" VALUES ('0987654321', '0987654321', 24);
INSERT INTO "public"."users" VALUES ('09876543', '0987654321', 25);
INSERT INTO "public"."users" VALUES ('112233', '112233', 26);
INSERT INTO "public"."users" VALUES ('1223', '1223', 27);
INSERT INTO "public"."users" VALUES ('1231', '1231', 28);
INSERT INTO "public"."users" VALUES ('1890', '1890', 29);
INSERT INTO "public"."users" VALUES ('0987', '0987', 30);

-- ----------------------------
-- Primary Key structure for table cards
-- ----------------------------
ALTER TABLE "public"."cards" ADD CONSTRAINT "cards_pkey" PRIMARY KEY ("name");

-- ----------------------------
-- Primary Key structure for table leaderboards
-- ----------------------------
ALTER TABLE "public"."leaderboards" ADD CONSTRAINT "leaderboards_pkey" PRIMARY KEY ("playerID");

-- ----------------------------
-- Primary Key structure for table monsters
-- ----------------------------
ALTER TABLE "public"."monsters" ADD CONSTRAINT "monsters_pkey" PRIMARY KEY ("name");

-- ----------------------------
-- Primary Key structure for table userconfigs
-- ----------------------------
ALTER TABLE "public"."userconfigs" ADD CONSTRAINT "userconfigs_pkey" PRIMARY KEY ("playerID");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("username");
