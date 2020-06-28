CREATE TABLE IF NOT EXISTS player_summary (
			steamid TEXT PRIMARY KEY NOT NULL,
            avatar TEXT,
            avatarfull TEXT,
            avatarmedium TEXT,
            cityid TEXT,
            commentpermission TEXT,
            communityvisibilitystate TEXT,
            gameextrainfo TEXT,
            gameid TEXT,
            gameserverip TEXT,
            lastlogoff TEXT,
            loccityid TEXT,
            loccountrycode TEXT,
            locstatecode TEXT,
            personaname TEXT,
            personastate TEXT,
            primaryclanid TEXT,
            profilestate TEXT,
            profileurl TEXT,
            realname TEXT,
            timecreated TEXT);

CREATE TABLE IF NOT EXISTS player_stats (
			steamid TEXT PRIMARY KEY NOT NULL,
			gi_lesson_bomb_sites_a TEXT,
			gi_lesson_bomb_sites_b TEXT,
			gi_lesson_csgo_cycle_weapons_kb TEXT,
			gi_lesson_csgo_hostage_lead_to_hrz TEXT,
			gi_lesson_csgo_instr_explain_bomb_carrier TEXT,
			gi_lesson_csgo_instr_explain_buyarmor TEXT,
			gi_lesson_csgo_instr_explain_buymenu TEXT,
			gi_lesson_csgo_instr_explain_follow_bomber TEXT,
			gi_lesson_csgo_instr_explain_inspect TEXT,
			gi_lesson_csgo_instr_explain_pickup_bomb TEXT,
			gi_lesson_csgo_instr_explain_plant_bomb TEXT,
			gi_lesson_csgo_instr_explain_prevent_bomb_pickup TEXT,
			gi_lesson_csgo_instr_explain_reload TEXT,
			gi_lesson_csgo_instr_explain_zoom TEXT,
			gi_lesson_csgo_instr_rescue_zone TEXT,
			gi_lesson_defuse_planted_bomb TEXT,
			gi_lesson_find_planted_bomb TEXT,
			gi_lesson_tr_explain_plant_bomb TEXT,
			gi_lesson_version_number TEXT,
			last_match_contribution_score TEXT,
			last_match_ct_wins TEXT,
			last_match_damage TEXT,
			last_match_deaths TEXT,
			last_match_dominations TEXT,
			last_match_favweapon_id TEXT,
			last_match_favweapon_hits TEXT,
			last_match_favweapon_kills TEXT,
			last_match_favweapon_shots TEXT,
			last_match_gg_contribution_score TEXT,
			last_match_kills TEXT,
			last_match_max_players TEXT,
			last_match_money_spent TEXT,
			last_match_mvps TEXT,
			last_match_revenges TEXT,
			last_match_rounds TEXT,
			last_match_t_wins TEXT,
			last_match_wins TEXT,
			steam_stat_matchwinscomp TEXT,
			steam_stat_survivedz TEXT,
			steam_stat_xpearnedgames TEXT,
			total_broken_windows TEXT,
			total_contribution_score TEXT,
			total_damage_done TEXT,
			total_deaths TEXT,
			total_defused_bombs TEXT,
			total_domination_overkills TEXT,
			total_dominations TEXT,
			total_gg_matches_played TEXT,
			total_gg_matches_won TEXT,
			total_gun_game_contribution_score TEXT,
			total_gun_game_rounds_played TEXT,
			total_gun_game_rounds_won TEXT,
			total_hits_ak47 TEXT,
			total_hits_aug TEXT,
			total_hits_awp TEXT,
			total_hits_bizon TEXT,
			total_hits_deagle TEXT,
			total_hits_elite TEXT,
			total_hits_famas TEXT,
			total_hits_fiveseven TEXT,
			total_hits_g3sg1 TEXT,
			total_hits_galilar TEXT,
			total_hits_glock TEXT,
			total_hits_hkp2000 TEXT,
			total_hits_m249 TEXT,
			total_hits_m4a1 TEXT,
			total_hits_mac10 TEXT,
			total_hits_mag7 TEXT,
			total_hits_mp7 TEXT,
			total_hits_mp9 TEXT,
			total_hits_negev TEXT,
			total_hits_nova TEXT,
			total_hits_p250 TEXT,
			total_hits_p90 TEXT,
			total_hits_sg556 TEXT,
			total_hits_sawedoff TEXT,
			total_hits_scar20 TEXT,
			total_hits_ssg08 TEXT,
			total_hits_tec9 TEXT,
			total_hits_ump45 TEXT,
			total_hits_xm1014 TEXT,
			total_kills TEXT,
			total_kills_against_zoomed_sniper TEXT,
			total_kills_ak47 TEXT,
			total_kills_aug TEXT,
			total_kills_awp TEXT,
			total_kills_bizon TEXT,
			total_kills_deagle TEXT,
			total_kills_elite TEXT,
			total_kills_enemy_blinded TEXT,
			total_kills_enemy_weapon TEXT,
			total_kills_famas TEXT,
			total_kills_fiveseven TEXT,
			total_kills_g3sg1 TEXT,
			total_kills_galilar TEXT,
			total_kills_glock TEXT,
			total_kills_headshot TEXT,
			total_kills_hegrenade TEXT,
			total_kills_hkp2000 TEXT,
			total_kills_knife TEXT,
			total_kills_knife_fight TEXT,
			total_kills_m249 TEXT,
			total_kills_m4a1 TEXT,
			total_kills_mac10 TEXT,
			total_kills_mag7 TEXT,
			total_kills_molotov TEXT,
			total_kills_mp7 TEXT,
			total_kills_mp9 TEXT,
			total_kills_negev TEXT,
			total_kills_nova TEXT,
			total_kills_p250 TEXT,
			total_kills_p90 TEXT,
			total_kills_sawedoff TEXT,
			total_kills_scar20 TEXT,
			total_kills_sg556 TEXT,
			total_kills_ssg08 TEXT,
			total_kills_taser TEXT,
			total_kills_tec9 TEXT,
			total_kills_ump45 TEXT,
			total_kills_xm1014 TEXT,
			total_matches_played TEXT,
			total_matches_won TEXT,
			total_matches_won_baggage TEXT,
			total_matches_won_bank TEXT,
			total_matches_won_lake TEXT,
			total_matches_won_safehouse TEXT,
			total_matches_won_shoots TEXT,
			total_matches_won_stmarc TEXT,
			total_matches_won_sugarcane TEXT,
			total_matches_won_train TEXT,
			total_money_earned TEXT,
			total_mvps TEXT,
			total_planted_bombs TEXT,
			total_progressive_matches_won TEXT,
			total_rescued_hostages TEXT,
			total_revenges TEXT,
			total_rounds_map_ar_baggage TEXT,
			total_rounds_map_ar_monastery TEXT,
			total_rounds_map_ar_shoots TEXT,
			total_rounds_map_cs_assault TEXT,
			total_rounds_map_cs_italy TEXT,
			total_rounds_map_cs_militia TEXT,
			total_rounds_map_cs_office TEXT,
			total_rounds_map_de_aztec TEXT,
			total_rounds_map_de_bank TEXT,
			total_rounds_map_de_cbble TEXT,
			total_rounds_map_de_dust TEXT,
			total_rounds_map_de_dust_2 TEXT,
			total_rounds_map_de_inferno TEXT,
			total_rounds_map_de_lake TEXT,
			total_rounds_map_de_nuke TEXT,
			total_rounds_map_de_safehouse TEXT,
			total_rounds_map_de_shorttrain TEXT,
			total_rounds_map_de_stmarc TEXT,
			total_rounds_map_de_sugarcane TEXT,
			total_rounds_map_de_train TEXT,
			total_rounds_map_de_vertigo TEXT,
			total_rounds_played TEXT,
			total_shots_ak47 TEXT,
			total_shots_aug TEXT,
			total_shots_awp TEXT,
			total_shots_bizon TEXT,
			total_shots_deagle TEXT,
			total_shots_elite TEXT,
			total_shots_famas TEXT,
			total_shots_fired TEXT,
			total_shots_fiveseven TEXT,
			total_shots_g3sg1 TEXT,
			total_shots_galilar TEXT,
			total_shots_glock TEXT,
			total_shots_hit TEXT,
			total_shots_hkp2000 TEXT,
			total_shots_m249 TEXT,
			total_shots_m4a1 TEXT,
			total_shots_mac10 TEXT,
			total_shots_mag7 TEXT,
			total_shots_mp7 TEXT,
			total_shots_mp9 TEXT,
			total_shots_negev TEXT,
			total_shots_nova TEXT,
			total_shots_p250 TEXT,
			total_shots_p90 TEXT,
			total_shots_sawedoff TEXT,
			total_shots_scar20 TEXT,
			total_shots_sg556 TEXT,
			total_shots_ssg08 TEXT,
			total_shots_taser TEXT,
			total_shots_tec9 TEXT,
			total_shots_ump45 TEXT,
			total_shots_xm1014 TEXT,
			total_time_played TEXT,
			total_trbomb_matches_won TEXT,
			total_tr_defused_bombs TEXT,
			total_tr_planted_bombs TEXT,
			total_weapons_donated TEXT,
			total_wins TEXT,
			total_wins_map_ar_baggage TEXT,
			total_wins_map_ar_monastery TEXT,
			total_wins_map_ar_shoots TEXT,
			total_wins_map_cs_assault TEXT,
			total_wins_map_cs_italy TEXT,
			total_wins_map_cs_militia TEXT,
			total_wins_map_cs_office TEXT,
			total_wins_map_de_aztec TEXT,
			total_wins_map_de_bank TEXT,
			total_wins_map_de_cbble TEXT,
			total_wins_map_de_dust TEXT,
			total_wins_map_de_dust_2 TEXT,
			total_wins_map_de_house TEXT,
			total_wins_map_de_inferno TEXT,
			total_wins_map_de_lake TEXT,
			total_wins_map_de_nuke TEXT,
			total_wins_map_de_safehouse TEXT,
			total_wins_map_de_shorttrain TEXT,
			total_wins_map_de_stmarc TEXT,
			total_wins_map_de_sugarcane TEXT,
			total_wins_map_de_train TEXT,
			total_wins_map_de_vertigo TEXT,
			total_wins_pistolround TEXT);

CREATE TABLE IF NOT EXISTS player_extra (
			steamid TEXT PRIMARY KEY NOT NULL,
			total_kd TEXT,
			last_match_kd TEXT,
			hit_ratio TEXT,
			played_hours TEXT,
			total_adr TEXT,
			last_match_adr TEXT);

CREATE TABLE IF NOT EXISTS recently_played (
			steamid TEXT PRIMARY KEY NOT NULL,
            appid TEXT,
            img_icon_url TEXT,
            img_logo_url TEXT,
            name TEXT,
            playtime_2_weeks TEXT,
            playtime_forever TEXT,
            playtime_linux_forever TEXT,
            playtime_mac_forever TEXT,
            playtime_windows_forever TEXT);

CREATE TABLE IF NOT EXISTS player_history (
			steamid TEXT NOT NULL,
			time DATETIME,
			total_kills TEXT,
			total_adr TEXT,
			total_shots_hit TEXT,
			total_shots_fired TEXT,
			total_kills_headshot TEXT,
			total_kd TEXT,
			last_match_contribution_score TEXT,
			last_match_damage TEXT,
			last_match_deaths TEXT,
			last_match_kills TEXT,
			last_match_rounds TEXT,
			last_match_kd TEXT,
			last_match_adr TEXT,
			hit_ratio TEXT,
			playtime_2_weeks TEXT);


