alter table security_cameras add constraint security_cameras_fieldofview_check CHECK ( field_of_view between 1 and 120);

alter table security_cameras add constraint security_cameras_storagecapacity_check check ( storage_capacity > 0 );

alter table security_cameras add constraint security_cameras_recordingduration_check check ( recording_duration > 0 );

