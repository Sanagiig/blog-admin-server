# 获取所有父节点 ID

# DELIMITER $$;
DROP PROCEDURE IF EXISTS getAllParentID;
CREATE DEFINER=`ginblog`@`%` PROCEDURE `getAllParentID`(IN id VARCHAR(255),IN tbName VARCHAR(255),OUT ids LONGTEXT)
BEGIN
    DECLARE counter INT DEFAULT 0;
    SET ids = id;
    SET @tmpID = id;
    SET @tmpParentID = id;

    SET @sqlStr = CONCAT('SELECT parent_id into @tmpParentID from ',tbName,' WHERE ',tbName,'.id = @tmpID');
    PREPARE findParent FROM @sqlStr;

    WHILE (@tmpID IS NOT NULL ) AND (@tmpID <> '') AND (counter < 100) DO
            EXECUTE findParent;
            SET @tmpID = @tmpParentID;
            SET counter = counter +1;

            IF (@tmpParentID IS NOT NULL)THEN
                SET ids = CONCAT(@tmpParentID,',',ids);
            END IF;
        END WHILE;
END;

# 获取所有子节点 ID
DROP PROCEDURE IF EXISTS getAllChildrenID;
CREATE DEFINER=`ginblog`@`%` PROCEDURE `getAllChildrenID`(IN id VARCHAR(255),IN tbName VARCHAR(255),OUT ids LONGTEXT)
BEGIN
    DECLARE counter INT DEFAULT 0;
    SET ids = id;
    SET @tmpID = id;
    SET @tmpChildrenIDs = id;

    SET @sqlStr = CONCAT('SELECT GROUP_CONCAT(id) into @tmpChildrenIDs from ',tbName,' WHERE FIND_IN_SET(',tbName,'.parent_id, @tmpChildrenIDs)');
    PREPARE findParent FROM @sqlStr;

    WHILE (@tmpChildrenIDs <> '') AND (counter < 100) DO
            EXECUTE findParent;

            SET counter = counter +1;
            SET ids = CONCAT_WS(',',ids,@tmpChildrenIDs);

        END WHILE;
END