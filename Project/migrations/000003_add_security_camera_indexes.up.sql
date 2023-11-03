CREATE INDEX IF NOT EXISTS security_camera_manufacturer_idx ON security_cameras USING GIN (to_tsvector('simple', manufacturer));
CREATE INDEX IF NOT EXISTS security_camera_resolution_idx ON security_cameras USING GIN (to_tsvector('simple', resolution));