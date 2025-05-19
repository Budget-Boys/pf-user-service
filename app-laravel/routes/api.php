<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\UserController;


Route::get('/ping', function () {
    return response()->json(['pong' => true]);
});

Route::get('/teste', function () {
    return response()->json(['ok' => true]);
});

Route::post('/user', [UserController::class, 'store']);