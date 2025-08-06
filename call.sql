call sp_nightly_call();

-- -- call after nightly ETL
-- call lg.sp_szn_load();
-- call lg.sp_team_all_load();
-- call lg.sp_plr_all_load();
-- call stats.sp_tbox();
-- call stats.sp_pbox();
-- call api.sp_plr_agg();